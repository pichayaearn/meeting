package route

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pichayaearn/meeting/pkg/auth/serializer"
	model "github.com/pichayaearn/meeting/pkg/model"
)

type CreateUserCfg struct {
	UserSvc model.UserSvc
}

func CreateUser(cfg CreateUserCfg) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(serializer.CreateUserReq)

		// Use BindJSON() to bind the request body as JSON into the user struct
		if err := c.Bind(req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body"+err.Error())
		}

		//validate request
		if err := req.ValidateRequest(); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body: "+err.Error())
		}

		err := cfg.UserSvc.CreateUser(model.CreateUser{
			Email:    req.Email,
			Password: req.Password,
		})
		if err != nil {
			return echo.NewHTTPError(http.StatusForbidden, "Create user"+err.Error())
		}

		return c.NoContent(http.StatusCreated)

	}
}
