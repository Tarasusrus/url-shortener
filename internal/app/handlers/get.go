package handlers

import (
	"github.com/Tarasusrus/url-shortener/internal/app/stores"
	"log"
	"net/http"
	"strings"
)

// HandleGet обрабатывает GET-запросы
func HandleGet(w http.ResponseWriter, r *http.Request, store *stores.Store) {
	log.Printf("Received request from: %s", r.RemoteAddr)
	id := strings.TrimPrefix(r.URL.Path, "/")
	log.Printf("Received ID: %s", id)
	if id == "" {
		log.Printf("Empty ID received, responding with BadRequest")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	url, ok := store.Get(id)
	log.Printf("Retrieved URL: %s, Found: %v", url, ok)

	if ok {
		w.Header().Set("Location", url)
		w.WriteHeader(http.StatusTemporaryRedirect)
		log.Printf("Redirecting to: %s", url)
		return
	}
	log.Printf("URL not found for ID: %s, Responding with BadRequest", id)
	w.WriteHeader(http.StatusBadRequest)
}
