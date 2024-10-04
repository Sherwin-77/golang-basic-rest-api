package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/sherwin-77/golang-basic-rest-api/requests"
	"github.com/sherwin-77/golang-basic-rest-api/resources"
	"github.com/sherwin-77/golang-basic-rest-api/services"
)

type TodoHandler struct {
	TodoService *services.TodoService
}

func (th *TodoHandler) GetTodos(ctx echo.Context) error {
	todos := th.TodoService.GetTodos()
	resource := resources.TodoResource{}
	return ctx.JSON(200, resource.Collections(todos))
}

func (th *TodoHandler) GetTodo(ctx echo.Context) error {
	id := ctx.Param("id")
	requests.ValidateUUID(id)

	todo := th.TodoService.GetTodoByID(id)
	resource := resources.TodoResource{}
	return ctx.JSON(200, resource.Make(todo))
}

func (th *TodoHandler) CreateTodo(ctx echo.Context) error {
	var todoRequest requests.TodoRequest

	if err := ctx.Bind(&todoRequest); err != nil {
		panic(err)
	}

	if err := ctx.Validate(&todoRequest); err != nil {
		panic(err)
	}

	todo := th.TodoService.CreateTodo(todoRequest)
	resource := resources.TodoResource{}
	return ctx.JSON(201, resource.Make(todo))
}

func (th *TodoHandler) UpdateTodo(ctx echo.Context) error {
	var todoRequest requests.TodoUpdateRequest
	id := ctx.Param("id")
	requests.ValidateUUID(id)

	if err := ctx.Bind(&todoRequest); err != nil {
		panic(err)
	}

	if err := ctx.Validate(&todoRequest); err != nil {
		panic(err)
	}

	todo := th.TodoService.UpdateTodoByID(id, todoRequest)
	resource := resources.TodoResource{}
	return ctx.JSON(200, resource.Make(todo))
}

func (th *TodoHandler) DeleteTodo(ctx echo.Context) error {
	id := ctx.Param("id")
	requests.ValidateUUID(id)

	th.TodoService.DeleteTodoByID(id)
	return ctx.NoContent(204)
}
