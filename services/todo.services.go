package services

import (
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sherwin-77/golang-basic-rest-api/db"
	"github.com/sherwin-77/golang-basic-rest-api/models"
	"github.com/sherwin-77/golang-basic-rest-api/requests"
)

type TodoService struct {
}

func (ts *TodoService) GetTodos() []models.Todo {
	conn := db.GetDB()
	// Get todos from database
	rows, err := conn.Query("SELECT * FROM todos")
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	defer db.ReleaseDB(conn)

	var todos = []models.Todo{}

	// Loop through all todos and append to todos slice
	for rows.Next() {
		var todo models.Todo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.IsCompleted, &todo.CreatedAt, &todo.UpdatedAt)
		if err != nil {
			panic(err)
		}
		todos = append(todos, todo)
	}

	if err = rows.Err(); err != nil {
		panic(err)
	}

	return todos
}

func (ts *TodoService) GetTodoByID(id string) models.Todo {
	conn := db.GetDB()
	// Get todo from database
	rows, err := conn.Query("SELECT * FROM todos WHERE id = $1 LIMIT 1", id)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	defer db.ReleaseDB(conn)

	var todo models.Todo

	// Check if todo exists. If not, return panic not found
	if rows.Next() {
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.IsCompleted, &todo.CreatedAt, &todo.UpdatedAt)
		if err != nil {
			panic(err)
		}
	} else {
		panic(echo.NewHTTPError(404, "Todo not found"))
	}

	if err = rows.Err(); err != nil {
		panic(err)
	}

	return todo
}

func (ts *TodoService) CreateTodo(todoRequest requests.TodoRequest) models.Todo {
	conn := db.GetDB()
	// Create new todo
	todo := models.Todo{
		ID:          uuid.Must(uuid.NewV7()).String(),
		Title:       todoRequest.Title,
		Description: todoRequest.Description,
		IsCompleted: todoRequest.IsCompleted,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Insert todo to database
	_, err := conn.Exec("INSERT INTO todos (id, title, description, is_completed, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)", todo.ID, todo.Title, todo.Description, todo.IsCompleted, todo.CreatedAt, todo.UpdatedAt)
	if err != nil {
		panic(err)
	}
	defer db.ReleaseDB(conn)

	return todo
}

func (ts *TodoService) UpdateTodoByID(id string, todoRequest requests.TodoUpdateRequest) models.Todo {
	conn := db.GetDB()
	todo := ts.GetTodoByID(id)

	// TODO: Maybe refactor
	if todoRequest.Title != todo.Title {
		todo.Title = todoRequest.Title
	}
	if todoRequest.Description != todo.Description {
		todo.Description = todoRequest.Description
	}
	if todoRequest.IsCompleted != todo.IsCompleted {
		todo.IsCompleted = todoRequest.IsCompleted
	}
	todo.UpdatedAt = time.Now()

	// Update todo in database
	_, err := conn.Exec("UPDATE todos SET title = $1, description = $2, is_completed = $3, updated_at = $4 WHERE id = $5", todo.Title, todo.Description, todo.IsCompleted, todo.UpdatedAt, todo.ID)
	if err != nil {
		panic(err)
	}

	defer db.ReleaseDB(conn)

	return todo
}

func (ts *TodoService) DeleteTodoByID(id string) {
	conn := db.GetDB()

	todo := ts.GetTodoByID(id)

	// Delete todo from database
	_, err := conn.Exec("DELETE FROM todos WHERE id = $1", todo.ID)
	if err != nil {
		panic(err)
	}

	defer db.ReleaseDB(conn)
}
