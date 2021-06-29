package main

import (
	"net/http"

	"github.com/fraidev/todo-go/internal/endpoints"
	"github.com/fraidev/todo-go/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	store := store.NewTodoStoreDB()
	todoEndpoints := endpoints.NewTodoEndpoints(store)
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Route("/todo", func(r chi.Router) {
		r.Get("/", todoEndpoints.GetAll)
		r.Post("/", todoEndpoints.Create)
		r.Get("/{id}", todoEndpoints.GetById)
		r.Put("/{id}", todoEndpoints.Update)
		r.Delete("/{id}", todoEndpoints.Delete)
	})
	http.ListenAndServe(":3000", r)
}
