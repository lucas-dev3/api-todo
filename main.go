package main

import (
	"api-postgres/configs"
	"api-postgres/handlers"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	// This is a placeholder for the main function

	err := configs.Load()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()

	r.Post("/todos", handlers.Create)
	r.Put("/todos/{id}", handlers.Update)
	r.Delete("/todos/{id}", handlers.Delete)
	r.Get("/todos", handlers.List)
	r.Get("/todos/{id}", handlers.GetById)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetAPIConfig()), r)
}
