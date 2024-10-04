package requests

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func ValidateUUID(id string) error {
	if _, err := uuid.Parse(id); err != nil {
		panic(echo.NewHTTPError(400, "Invalid ID format"))
	}
	return nil
}
