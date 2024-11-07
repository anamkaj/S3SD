package stat_client

import (
	"direct/internal/models"

	"net/http"
	"net/http/httptest"

	"testing"
)

func TestHandler_RegisterRoutes(t *testing.T) {

	statStore := &mockStatStore{}
	handler := NewHandler(statStore)

	t.Run("handleStatistics", func(t *testing.T) {

		req, err := http.NewRequest(http.MethodGet, "/api/client_stat?client_login=test", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		router := http.NewServeMux()

		router.HandleFunc("/api/client_stat", handler.handleStatistics)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				rr.Code, http.StatusOK)
		}
	})
}

type mockStatStore struct{}

func (m *mockStatStore) GetStatClient(clientLogin string) (*[]models.StatisticsClient, error) {

	return nil, nil
}
