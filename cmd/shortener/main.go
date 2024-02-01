package main

import (
	"github.com/Tarasusrus/url-shortener/internal/app"
	"io"
	"mime"
	"net/http"
	"strings"
)

// handlePost обрабатывает POST-запросы
func handlePost(w http.ResponseWriter, r *http.Request, store *app.Store) {
	// Проверяем, что ContentType запроса text/plain
	// Если это не так, то возвращает код ошибки 400
	mediaType, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil || mediaType != "text/plain" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Преобразуем тело запроса в строку, сокращаем URL и сохраняем его
	url := string(body)
	id := store.Set(url)

	// Отправляем ответ с кодом 201 и сокращенным URL
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(id))
}

// handleGet обрабатывает GET-запросы
func handleGet(w http.ResponseWriter, r *http.Request, store *app.Store) {
	id := strings.TrimPrefix(r.URL.Path, "/")
	url, ok := store.Get(id)
	if ok {
		w.Header().Set("Location", url)
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}

func main() {
	store := app.NewStore()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			handlePost(w, r, store)
		case http.MethodGet:
			handleGet(w, r, store)
		default:
			// На любой некорректный запрос сервер должен возвращать ответ с кодом 400
			w.WriteHeader(http.StatusBadRequest)
		}
	})
	http.ListenAndServe(":8080", nil)
}
