package handler

import (
	"context"
	"net/http"

	"github.com/labstack/echo"

	"echo-demo/internal/architecture/three/domain"
	"echo-demo/internal/architecture/three/infrastructure"
)

type TodoHnadler struct {
	todoRepo *infrastructure.TodoRepository
}

func NewTodoHandler(todoRepo *infrastructure.TodoRepository) *TodoHnadler{
	return &TodoHnadler{todoRepo: todoRepo}
}

func(h *TodoHnadler) ListTodos(c echo.Context) error {
	ctx := context.Background()

	todos, err := h.todoRepo.ListTodos(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, todos)
}

func (h *TodoHnadler) RegisterTodo(c echo.Context) error {
	var todo domain.Todo
	if err := c.Bind(&todo); err != nil{
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := context.Background()
	err := h.todoRepo.RegisterTodo(ctx, &todo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, todo.ID)
}
