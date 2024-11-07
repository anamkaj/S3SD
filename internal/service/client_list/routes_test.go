package list_client

import (
	"direct/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler_RegisterRout(t *testing.T) {

	listStore := &mockListStore{}
	handler := NewHandler(listStore)

	t.Run("handleListClients", func(t *testing.T) {

		req, err := http.NewRequest(http.MethodGet, "/api/clients_list", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := http.NewServeMux()

		router.HandleFunc("/api/clients_list", handler.handleListClients)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				rr.Code, http.StatusOK)
		}

	})
}

type mockListStore struct{}

func (m *mockListStore) GetClientList() (*[]models.List, error) {
	return nil, nil
}

func (m *mockListStore) InsertClientList(list *models.ResApiDirect) error {
	return nil
}
