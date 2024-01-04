package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	chi "github.com/go-chi/chi/v5"

	"app/internal/handler"
)

func main() {
	h := handler.NewHandler()

	r := chi.NewRouter()
	r.Get("/health", h.Health)

	addr := fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))
	err := http.ListenAndServe(addr, r)
	log.Fatal(err)
}
