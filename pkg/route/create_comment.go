package route

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/pichayaearn/meeting/pkg/model"
	"github.com/pichayaearn/meeting/pkg/serializer"
)

type CreateCommentCfg struct {
	CommentSvc model.CommentSvc
}

func CreateComment(cfg CreateCommentCfg) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID, err := serializer.BindUserIDFromContext(c)
		if err != nil {
			return echo.NewHTTPError(http.StatusForbidden, "Bind user id "+err.Error())
		}

		req := serializer.NewCreateCommentReq(userID)

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

		if err := cfg.CommentSvc.Create(model.CreateCommentOpts{
			MeetingID: meetingID,
			Detail:    req.Detail,
			CreatedBy: req.CreatedBy,
		}); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Create  "+err.Error())
		}

		return c.NoContent(http.StatusCreated)
	}
}
