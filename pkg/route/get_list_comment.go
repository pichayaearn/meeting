package route

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/pichayaearn/meeting/pkg/model"
	"github.com/pichayaearn/meeting/pkg/serializer"
)

type GetListCommentCfg struct {
	CommentSvc model.CommentSvc
}

func GetListComment(cfg GetListCommentCfg) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(serializer.GetListCommentReq)

		// Use BindJSON() to bind the request body as JSON into the user struct
		if err := c.Bind(req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body"+err.Error())
		}

		//validate request
		if err := req.Validate(); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body: "+err.Error())
		}

		meetingID, err := uuid.Parse(req.MeetingID)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body: "+err.Error())
		}
		meetings, err := cfg.CommentSvc.List(model.GetListCommentOpts{
			MeetingID: meetingID,
			Limit:     req.Limit,
			Offset:    req.Offset,
		}, c.Request().Context())
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Get list meeting: "+err.Error())
		}

		resp := []serializer.GetListCommentResponse{}
		for _, v := range meetings {
			resp = append(resp, serializer.ToGetListCommentResponse(v))

		}

		return c.JSON(http.StatusOK, resp)

	}
}
