package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Seiji-Ikeda32/go-todo-app/backend/models"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	healthCheck := &models.HealthCheck{
		Message: "OK",
	}
	m, _ := json.Marshal(healthCheck)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Write(m)

}
