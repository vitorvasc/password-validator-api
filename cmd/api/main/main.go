package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/vitorvasc/api-password-validator/internal/api/handlers"
	"github.com/vitorvasc/api-password-validator/internal/domain/password"
)

func main() {
	// Initialize the router
	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Create validator and handler
	validator := password.NewValidator()
	passwordHandler := handlers.NewPasswordHandler(validator)

	// Routes
	r.Route("/v1", func(r chi.Router) {
		r.Post("/users/validate-password", passwordHandler.ValidatePassword)
	})

	// Start server
	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
