package handlers

import (
	"github.com/Tarasusrus/url-shortener/internal/app/configs"
	"github.com/Tarasusrus/url-shortener/internal/app/stores"
	"io"
	"log"
	"mime"
	"net/http"
)

// HandlePost обрабатывает POST-запросы
func HandlePost(w http.ResponseWriter, r *http.Request, store *stores.Store) {
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

	//scheme := "http"
	//if r.TLS != nil {
	//	scheme = "https"
	//}
	config := configs.NewFlagConfig()
	baseURL := config.BaseURL()
	shortURL := baseURL + id
	// Отправляем ответ с кодом 201 и сокращенным URL
	w.WriteHeader(http.StatusCreated)
	//urlPath := fmt.Sprintf("%s://%s/%s", scheme, r.Host, id)
	w.Write([]byte(shortURL))
}
