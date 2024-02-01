package main

import (
	"github.com/Tarasusrus/url-shortener/internal/app"
	"io"
	"log"
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
		log.Printf("Error parsing media type or media type is not text/plain: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Printf("Error reading request body: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Println(len(body))
	if len(body) == 0 {
		log.Printf("Received empty request body\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Преобразуем тело запроса в строку, сокращаем URL и сохраняем его
	url := string(body)
	id := store.Set(url)
	log.Printf("Short URL created: %s\n", id)

	// Отправляем ответ с кодом 201 и сокращенным URL
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(id))
}

// handleGet обрабатывает GET-запросы
func handleGet(w http.ResponseWriter, r *http.Request, store *app.Store) {
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

func main() {
	store := app.NewStore()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handleGet(w, r, store)
		case http.MethodPost:
			handlePost(w, r, store)
			//default:
			//	// На любой некорректный запрос сервер должен возвращать ответ с кодом 400
			//	w.WriteHeader(http.StatusBadRequest)
		}
	})
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
