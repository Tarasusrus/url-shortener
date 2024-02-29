package handlers

import (
	"log"
	"net/http"
	"strings"

	"github.com/Tarasusrus/url-shortener/internal/app/stores"
)

// HandleGet обрабатывает GET-запросы.
func HandleGet(responseWriter http.ResponseWriter, request *http.Request, store *stores.Store) {
	log.Printf("Received request from: %s", request.RemoteAddr)
	shortURLID := strings.TrimPrefix(request.URL.Path, "/")
	log.Printf("Received ID: %s", shortURLID)

	if shortURLID == "" {
		log.Printf("Empty ID received, responding with BadRequest")
		responseWriter.WriteHeader(http.StatusBadRequest)

		return
	}

	url, ok := store.Get(shortURLID)
	log.Printf("Retrieved URL: %s, Found: %v", url, ok)

	if !ok {
		log.Printf("URL not found for ID: %s, Responding with BadRequest", shortURLID)
		responseWriter.WriteHeader(http.StatusBadRequest)

		return
	}

	responseWriter.Header().Set("Location", url)
	responseWriter.WriteHeader(http.StatusTemporaryRedirect)
	log.Printf("Redirecting to: %s", url)
}
