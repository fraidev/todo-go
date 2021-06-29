package store

import (
	"errors"

	"github.com/fraidev/todo-go/internal/entities"
	"github.com/google/uuid"
)

type todoStoreMemory struct {
}

var todos = []entities.Todo{
	{ID: uuid.New(), Name: "todo", Completed: true},
}

func NewTodoStoreMemory() TodoStore {
	return &todoStoreMemory{}
}

func (t todoStoreMemory) GetById(id uuid.UUID) (*entities.Todo, error) {
	for _, i := range todos {
		if i.ID == id {
			return &i, nil
		}
	}

	return nil, errors.New("todo not found")
}

func (t todoStoreMemory) List() ([]entities.Todo, error) {
	return todos, nil
}

func (t todoStoreMemory) Create(todo *entities.Todo) error {
	todos = append(todos, *todo)
	return nil
}

func (t todoStoreMemory) Update(id uuid.UUID, todo *entities.Todo) error {
	for i := range todos {
		if todos[i].ID == id {
			todos[i] = *todo
			return nil
		}
	}

	return errors.New("todo not found")
}

func (t todoStoreMemory) DeleteById(id uuid.UUID) error {
	for index, i := range todos {
		if i.ID == id {
			todos = append(todos[:index], todos[index+1:]...)
			return nil
		}
	}

	return errors.New("todo not found")
}
