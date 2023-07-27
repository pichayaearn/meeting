package route

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/pichayaearn/meeting/pkg/model"
	"github.com/pichayaearn/meeting/pkg/serializer"
)

type GetListMeetingCfg struct {
	MeetingSvc model.MeetingSvc
}

func GetListMeeting(cfg GetListMeetingCfg) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(serializer.GetListMeetingReq)

		// Use BindJSON() to bind the request body as JSON into the user struct
		if err := c.Bind(req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body"+err.Error())
		}

		//validate request
		if err := req.Validate(); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body: "+err.Error())
		}

		var idUUID, createdByUUID uuid.UUID

		if req.ID != "" {
			id, err := uuid.Parse(req.ID)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body: "+err.Error())
			}
			idUUID = id
		}

		if req.CreatedBy != "" {
			createdBy, err := uuid.Parse(req.CreatedBy)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body: "+err.Error())
			}
			createdByUUID = createdBy
		}
		meetings, err := cfg.MeetingSvc.List(model.GetMeetingOpts{
			ID:        idUUID,
			CreatedBy: createdByUUID,
			Status:    req.Status,
			Limit:     req.Limit,
			Offset:    req.Offset,
		}, c.Request().Context())
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Get list meeting: "+err.Error())
		}

		if len(meetings) <= 0 {
			return echo.NewHTTPError(http.StatusNotFound, "meeting not found")
		}

		resp := []serializer.GetListMeetingResponse{}

		for _, v := range meetings {
			meeting := serializer.ToGetListMeetingResponse(v)
			resp = append(resp, meeting)
		}

		return c.JSON(http.StatusOK, resp)

	}
}
