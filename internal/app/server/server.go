package server

import (
	"github.com/Tarasusrus/url-shortener/internal/app/handlers"
	"github.com/Tarasusrus/url-shortener/internal/app/stores"
	"log"
	"net/http"
)

func Run() {
	store := stores.NewStore()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.HandleGet(w, r, store)
		case http.MethodPost:
			handlers.HandlePost(w, r, store)
		default:
			w.WriteHeader(http.StatusBadRequest) // На любой некорректный запрос сервер должен возвращать ответ с кодом 400
		}
	})
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
