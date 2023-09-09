package domain

import (
	"errors"
	"time"
)

type Todo struct {
	ID        int
	Title     string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewTodo(title, content string) (*Todo, error) {
	if len(title) == 0 || len(title) > 100 {
		return nil, errors.New("title must be between 1 and 100 characters")
	}
	if len(content) > 100 {
		return nil, errors.New("content must be less than or equal to 100 characters")
	}

	return &Todo{
		Title: title,
		Content: content,
	}, nil
}