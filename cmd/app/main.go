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

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	if host == "" {
		host = "0.0.0.0"
	}
	if port == "" {
		port = "8080"
	}
	addr := fmt.Sprintf("%s:%s", host, port)
	fmt.Println("Starting server", addr, "...")
	err := http.ListenAndServe(addr, r)
	log.Fatal(err)
}
