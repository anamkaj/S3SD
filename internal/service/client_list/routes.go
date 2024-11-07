package list_client

import (
	"direct/internal/models"
	"direct/internal/request"
	"direct/internal/utils"

	"fmt"
	"net/http"
)

type Handler struct {
	store models.ListClientStore
}

func NewHandler(store models.ListClientStore) *Handler {
	return &Handler{
		store: store,
	}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.Handle("GET /api/clients_list", http.HandlerFunc(h.handleListClients))
}

func (h *Handler) handleListClients(w http.ResponseWriter, r *http.Request) {

	responseBody, err := request.GetAgencyClients()
	if err != nil {
		http.Error(w, "Ошибка при получении списка клиентов", http.StatusInternalServerError)
		return
	}

	err = h.store.InsertClientList(responseBody)
	if err != nil {
		fmt.Println("Ошибка добавления клиентов:", err)
	}

	fmt.Println("Список клиентов обновлен")

	data, err := h.store.GetClientList()
	if err != nil {
		fmt.Println("Ошибка чтения списка  клиентов:", err)
	}

	utils.ResJson(w, http.StatusOK, data)
}
