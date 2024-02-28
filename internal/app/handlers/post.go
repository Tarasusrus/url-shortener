// Package handlers предоставляет обработчики для HTTP-запросов.
// В этом пакете определены функции для обработки входящих запросов и отправки соответствующих ответов.
package handlers

import (
	"fmt"
	"io"
	"mime"
	"net/http"

	"github.com/Tarasusrus/url-shortener/internal/app/configs"
	"github.com/Tarasusrus/url-shortener/internal/app/stores"
	"github.com/Tarasusrus/url-shortener/internal/logger"
	"go.uber.org/zap"
)

// HandlePost обрабатывает POST-запросы на сокращение URL.
// Эта функция читает тело запроса как текст/plain и использует хранилище для генерации и сохранения сокращенного URL.
// В случае успеха, клиенту возвращается статус 201 (Created) и сокращенный URL.
// Если запрос не удовлетворяет требованиям (например, неправильный ContentType или пустое тело запроса),
// клиенту возвращается статус 400 (BadRequest).
func HandlePost(
	responseWriter http.ResponseWriter, request *http.Request, store *stores.Store, config *configs.FlagConfig) {
	// Проверка ContentType на соответствие text/plain.
	mediaType, _, err := mime.ParseMediaType(request.Header.Get("Content-Type"))

	//todo вынести в функцию повторяющийся код.
	if err != nil || mediaType != "text/plain" {
		logger.Log.Info("Invalid content type", zap.Error(err))
		responseWriter.WriteHeader(http.StatusBadRequest)
		return
	}

	// Чтение тела запроса.
	body, err := io.ReadAll(request.Body)

	if err != nil {
		logger.Log.Info("Error reading request body", zap.Error(err))
		responseWriter.WriteHeader(http.StatusBadRequest)
		return
	}

	defer func() {
		if err := request.Body.Close(); err != nil {
			logger.Log.Error("Failed to close request body", zap.Error(err))
		}
	}()

	// Обработка пустого тела запроса.
	if len(body) == 0 {
		logger.Log.Info("Received empty request body")
		responseWriter.WriteHeader(http.StatusBadRequest)
		return
	}

	// Сокращение URL и его сохранение.
	url := string(body)
	shortURLId := store.Set(url)
	logger.Log.Info("Short URL created", zap.String("shortURLId", shortURLId))

	// Формирование и отправка сокращенного URL клиенту.
	scheme := "http"

	if request.TLS != nil {
		scheme = "https"
	}

	host := config.GetAddress()
	urlPath := fmt.Sprintf("%s://%s/%s", scheme, host, shortURLId)

	responseWriter.WriteHeader(http.StatusCreated)

	if _, err := responseWriter.Write([]byte(urlPath)); err != nil {
		logger.Log.Error("Failed to write response", zap.Error(err))
	}
}
