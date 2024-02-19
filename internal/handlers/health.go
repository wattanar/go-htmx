package handlers

import (
	"go-htmx/internal/pkg/response"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, r, map[string]interface{}{
		"status": http.StatusOK,
	})
}
