package route

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pichayaearn/meeting/pkg/model"
	"github.com/pichayaearn/meeting/pkg/serializer"
)

type CreateMeetingCfg struct {
	MeetingSvc model.MeetingSvc
}

func CreateMeeting(cfg CreateMeetingCfg) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, err := serializer.BindUserIDFromContext(c)
		if err != nil {
			return echo.NewHTTPError(http.StatusForbidden, "Bind user id "+err.Error())
		}

		req := serializer.NewCreateMeetingReq(userID)

		// Use BindJSON() to bind the request body as JSON into the user struct
		if err := c.Bind(req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body"+err.Error())
		}

		//validate request
		if err := req.Validate(); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body: "+err.Error())
		}

		//create meeting
		if err := cfg.MeetingSvc.Create(model.CreateMeetingOpts{
			Title:     req.Title,
			Detail:    req.Detail,
			CreatedBy: req.CreatedBy,
		}); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Create  "+err.Error())
		}

		return c.NoContent(http.StatusCreated)
	}
}
