package main

import (
	"log"
	"net/http"

	chi "github.com/go-chi/chi/v5"

	"app/internal/handler"
)

func main() {
	h := handler.NewHandler()

	r := chi.NewRouter()
	r.Get("/health", h.Health)

	err := http.ListenAndServe(":8080", r)
	log.Fatal(err)
}
