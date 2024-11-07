package stat_client

import (
	"direct/internal/models"
	"encoding/json"
	"net/http"
)

type Handler struct {
	store models.StatStore
}

func NewHandler(store models.StatStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.Handle("GET /api/client_stat", http.HandlerFunc(h.handleStatistics))
}

// https://www.youtube.com/watch?v=7VLmLOiQ3ck

func (h *Handler) handleStatistics(w http.ResponseWriter, r *http.Request) {

	clientLogin := r.URL.Query().Get("client_login")

	if clientLogin == "" {
		http.Error(w, "Не указан логин клиента", http.StatusBadRequest)
		return
	}

	data, err := h.store.GetStatClient(clientLogin)
	if err != nil {
		http.Error(w, "Ошибка при получении статистики", http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Ошибка при кодировании в JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

}
