package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gorm.io/gorm"

	"echo-demo/internal/architecture/mono"
	"echo-demo/internal/architecture/three/handler"
	"echo-demo/internal/architecture/three/infrastructure"
)

func SetupRouter(db *gorm.DB) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	architectureGroup := e.Group("/architecture")
	architectureGroup.GET("/mono/todos", mono.ListTodos)
	architectureGroup.POST("/mono/todos", mono.RegisterTodo)

	todoRepo := infrastructure.NewTodoRepository(db)
	todoHandler := handler.NewTodoHandler(todoRepo)

	architectureGroup.GET("/three/todos", todoHandler.ListTodos)
	architectureGroup.POST("/three/todos", todoHandler.RegisterTodo)

	return e
}
