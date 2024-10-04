package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/sherwin-77/golang-basic-rest-api/handlers"
)

func RegisterRoutes(e *echo.Group) {
	todoHandler := &handlers.TodoHandler{}

	e.GET("/todos", todoHandler.GetTodos)
	e.GET("/todos/:id", todoHandler.GetTodo)
	e.POST("/todos", todoHandler.CreateTodo)
	e.PATCH("/todos/:id", todoHandler.UpdateTodo)
	e.DELETE("/todos/:id", todoHandler.DeleteTodo)
}
