package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Seiji-Ikeda32/go-todo-app/backend/models"
)

//　構造体を作成
type HealthHandler struct{}

//　関数が呼ばれた時に構造体を返す
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

//　json形式に変換して”ok”を返す
func (h *HealthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	healthCheck := &models.HealthCheck{
		Message: "OK",
	}
	m, _ := json.Marshal(healthCheck)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Write(m)

}
