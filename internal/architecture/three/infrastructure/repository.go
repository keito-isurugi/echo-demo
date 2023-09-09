package infrastructure

import (
	"context"
	"echo-demo/internal/architecture/three/domain"

	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{db: db}
}

func (r *TodoRepository) ListTodos(ctx context.Context) ([]*domain.Todo, error) {
	todos := []*domain.Todo{}
	err := r.db.Find(&todos).Error
	return todos, err
}

func (r *TodoRepository) RegisterTodo(ctx context.Context, todo *domain.Todo) error {
	return r.db.Create(todo).Error
}