package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sherwin-77/golang-basic-rest-api/configs"
	"github.com/sherwin-77/golang-basic-rest-api/db"
	"github.com/sherwin-77/golang-basic-rest-api/routes"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func customHTTPErrorHandler(err error, c echo.Context) {
	var code int
	var message interface{}

	// Specific error handling log require installing the gommon package
	// Check: https://github.com/labstack/echo/issues/1017

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		message = he.Message
	} else if ve, ok := err.(validator.ValidationErrors); ok {
		code = http.StatusUnprocessableEntity
		fieldErr := ve[0]
		switch fieldErr.Tag() {
		case "required":
			message = fieldErr.Field() + " is required"
		case "email":
			message = fieldErr.Field() + " is not a valid email"
		case "gte":
			message = fieldErr.Field() + " must be greater than or equal to " + fieldErr.Param()
		case "lte":
			message = fieldErr.Field() + " must be less than or equal to " + fieldErr.Param()
		default:
			message = fieldErr.Field() + " is not valid"
		}
	} else {
		code = http.StatusInternalServerError
		message = http.StatusText(http.StatusInternalServerError)
		c.Logger().Error(err.Error())
	}

	if !c.Response().Committed {
		c.JSON(code, map[string]interface{}{
			"error": message,
		})
	}
}

func main() {
	e := echo.New()
	e.HideBanner = true
	config := configs.GetConfiguration()
	db.InitDB(config.Database.Path)

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} ${method} ${uri} ${status}\n",
	}))
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		DisablePrintStack: true,
	}))
	e.Validator = &CustomValidator{validator: validator.New()}
	e.HTTPErrorHandler = customHTTPErrorHandler

	group := e.Group("/api")
	group.GET("", func(ctx echo.Context) error {
		return ctx.JSON(200, map[string]interface{}{
			"message": "Hello world!",
			"data":    nil,
		})
	})
	routes.RegisterRoutes(group)

	if err := e.Start(fmt.Sprintf("localhost:%d", config.Server.Port)); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
