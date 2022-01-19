package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../../models"
)

type HealthHandler struct {}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	healthCheck := &models.HealthCheck{
		Message: "OK",
	}
	m, _ := json.Marshal(healthCheck)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(m))
}