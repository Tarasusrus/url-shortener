package main

import (
	"github.com/Tarasusrus/url-shortener/internal/app"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	store := app.NewStore()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			// Проверяем, что ContentType запроса text/plain
			// Если это не так, то возвращает код ошибки 400
			if r.Header.Get("Content-Type") != "text/plain" {
				http.Error(w, "Invalid Content-Type, expected text/plain", http.StatusBadRequest)
				return
			}

			// Читаем тело запроса
			body, err := ioutil.ReadAll(r.Body)
			defer r.Body.Close()
			if err != nil {
				http.Error(w, "Error reading request body", http.StatusInternalServerError)
			}

			// Преобразуем тело запроса в строку, сокращаем URL и сохраняем его
			url := string(body)
			id := store.Set(url)

			// Отправляем ответ с кодом 201 и сокращенным URL
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(id))

		case http.MethodGet:
			id := strings.TrimPrefix(r.URL.Path, "/")
			url, ok := store.Get(id)
			if ok {
				w.Header().Set("Location", url)
				w.WriteHeader(http.StatusTemporaryRedirect)
				return
			}
			http.Error(w, "URL not found", http.StatusNotFound)

		default:
			// На любой некорректный запрос сервер должен возвращать ответ с кодом 400
			http.Error(w, "Invalid request method", http.StatusBadRequest)
		}
	})

	http.ListenAndServe(":8080", nil)
}
