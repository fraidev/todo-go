package store

import (
	"github.com/fraidev/todo-go/internal/entities"
	"github.com/google/uuid"
)

type TodoStore interface {
	GetById(id uuid.UUID) (*entities.Todo, error)
	List() ([]entities.Todo, error)
	Create(todo *entities.Todo) error
	Update(id uuid.UUID, todo *entities.Todo) error
	DeleteById(id uuid.UUID) error
}
