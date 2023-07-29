package route

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/pichayaearn/meeting/pkg/model"
	"github.com/pichayaearn/meeting/pkg/serializer"
)

type UpdateStatusMeetingCfg struct {
	MeetingSvc model.MeetingSvc
}

func UpdateStatusMeeting(cfg UpdateStatusMeetingCfg) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(serializer.UpdateMeetingReq)

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

		if err := cfg.MeetingSvc.Update(model.UpdateMeetingOpts{
			MeetingID: meetingID,
			Status:    req.Status,
		}); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Get list meeting: "+err.Error())
		}

		return c.NoContent(http.StatusOK)

	}
}
