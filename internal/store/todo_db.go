package store

import (
	"context"
	"log"
	"os"

	"github.com/fraidev/todo-go/internal/entities"
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
)

type todoStoreDB struct {
	db *pg.DB
}

func NewTodoStoreDB() TodoStore {
	cs := os.Getenv("POSTGRES_URL")
	opt, err := pg.ParseURL(cs)
	if err != nil {
		log.Println(err.Error())
	}

	db := pg.Connect(opt)
	return &todoStoreDB{db}
}

func (t todoStoreDB) GetById(id uuid.UUID) (*entities.Todo, error) {
	res := entities.Todo{}
	err := t.db.Model(&res).Where("id = ?", id.String()).Select()
	return &res, err
}

func (t todoStoreDB) List() ([]entities.Todo, error) {
	var res []entities.Todo
	err := t.db.Model(&res).Select()
	return res, err
}

func (t todoStoreDB) Create(todo *entities.Todo) error {
	_, err := t.db.Model(todo).Insert()
	return err
}

func (t todoStoreDB) Update(id uuid.UUID, todo *entities.Todo) error {
	_, err := t.db.Model(todo).WherePK().Update()
	return err
}

func (t todoStoreDB) DeleteById(id uuid.UUID) error {
	ctx := context.Background()
	_, err := t.db.
		Model((*entities.Todo)(nil)).Where("id = (?)", id.String()).
		Delete(ctx)
	return err
}
