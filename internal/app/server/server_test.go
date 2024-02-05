package server

import (
	"bytes"
	"github.com/Tarasusrus/url-shortener/internal/app/handlers"
	"github.com/Tarasusrus/url-shortener/internal/app/stores"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleGet(t *testing.T) {
	store := stores.NewStore() // Хранилище для теста
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { handlers.HandleGet(w, r, store) }))
	defer server.Close()

	res, err := http.Get(server.URL)
	if err != nil {
		t.Errorf("Failed to send GET request to server: %v", err)
	}
	if res.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status code BadRequest, got : %v", res.StatusCode)
	}
}

func TestHandlePost(t *testing.T) {
	store := stores.NewStore() // Хранилище для теста
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { handlers.HandlePost(w, r, store) }))
	defer server.Close()

	reqBody := "http://example.com"
	res, err := http.Post(server.URL, "text/plain", bytes.NewBufferString(reqBody))
	if err != nil {
		t.Errorf("Failed to send POST request to server: %v", err)
	}
	if res.StatusCode != http.StatusCreated {
		t.Errorf("Expected status code Created, got : %v", res.StatusCode)
	}
}
