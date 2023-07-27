package serializer

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func BindUserIDFromContext(c echo.Context) (uuid.UUID, error) {
	userID := c.Get("ActionBy")
	userIDStr, ok := userID.(string)
	if !ok {
		return uuid.Nil, fmt.Errorf("user id not string")
	}

	uid, err := uuid.Parse(userIDStr)
	if err != nil {
		return uuid.Nil, fmt.Errorf("parse user id to uuid")
	}
	return uid, nil
}
