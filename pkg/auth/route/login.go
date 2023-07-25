package route

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pichayaearn/meeting/pkg/auth/model"
	"github.com/pichayaearn/meeting/pkg/auth/serializer"
)

type LoginCfg struct {
	AuthSvc model.AuthSvc
}

func Login(cfg LoginCfg) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(serializer.LoginReq)

		// Use BindJSON() to bind the request body as JSON into the user struct
		if err := c.Bind(req); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body"+err.Error())
		}

		//validate request
		if err := req.ValidateRequest(); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body: "+err.Error())
		}

		token, err := cfg.AuthSvc.Login(req.Email, req.Password)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Login failed: "+err.Error())
		}

		resp := serializer.ToLoginResponse(token)

		return c.JSON(200, resp)
	}
}
