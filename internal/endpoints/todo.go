package endpoints

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/fraidev/todo-go/internal/entities"
	"github.com/fraidev/todo-go/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type TodoEndpoints struct {
	store store.TodoStore
}

func NewTodoEndpoints(store store.TodoStore) TodoEndpoints {
	return TodoEndpoints{store: store}
}

func (t TodoEndpoints) GetById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	uuid, _ := uuid.Parse(id)
	todo, _ := t.store.GetById(uuid)
	json.NewEncoder(w).Encode(todo)
}

func (t TodoEndpoints) GetAll(w http.ResponseWriter, r *http.Request) {
	todos, err := t.store.List()
	if err != nil {
		log.Println(err.Error())
	}
	json.NewEncoder(w).Encode(todos)
}

func (t TodoEndpoints) Create(w http.ResponseWriter, r *http.Request) {
	var todo entities.Todo
	json.NewDecoder(r.Body).Decode(&todo)
	t.store.Create(&todo)
}

func (t TodoEndpoints) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	uuid, _ := uuid.Parse(id)
	var todo entities.Todo
	json.NewDecoder(r.Body).Decode(&todo)
	t.store.Update(uuid, &todo)
}

func (t TodoEndpoints) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	uuid, _ := uuid.Parse(id)
	t.store.DeleteById(uuid)
}
