package handlers

import (
	"fmt"
	"github.com/Tarasusrus/url-shortener/helpers"
	"github.com/Tarasusrus/url-shortener/internal/app/configs"
	"github.com/Tarasusrus/url-shortener/internal/app/stores"
	"io"
	"log"
	"mime"
	"net/http"
)

// HandlePost обрабатывает POST-запросы
func HandlePost(w http.ResponseWriter, r *http.Request, store *stores.Store, config *configs.FlagConfig) {
	// Проверяем, что ContentType запроса text/plain
	// Если это не так, то возвращает код ошибки 400
	mediaType, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil || mediaType != "text/plain" {
		helpers.LogError(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		helpers.LogError(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
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

	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	host := config.GetAddress()

	// Отправляем ответ с кодом 201 и сокращенным URL
	w.WriteHeader(http.StatusCreated)
	urlPath := fmt.Sprintf("%s://%s/%s", scheme, host, id)
	w.Write([]byte(urlPath))
}
