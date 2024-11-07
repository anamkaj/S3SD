package stat_client

import (
	"direct/internal/models"
	"direct/internal/utils"
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

func (h *Handler) handleStatistics(w http.ResponseWriter, r *http.Request) {

	clientLogin := r.URL.Query().Get("client_login")

	if clientLogin == "" {
		http.Error(w, "не указан логин клиента", http.StatusBadRequest)
		return
	}

	data, err := h.store.GetStatClient(clientLogin)
	if err != nil {
		http.Error(w, "Ошибка при получении статистики", http.StatusInternalServerError)
		return
	}

	utils.ResJson(w, http.StatusOK, data)

}
