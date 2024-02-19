package routes

import (
	"go-htmx/internal/handlers"
	"net/http"
)

func NewRouter(r *http.ServeMux) {
	r.HandleFunc("GET /", handlers.HealthCheck)
}
