package handlers

import (
	"github.com/Tarasusrus/url-shortener/internal/app/configs"
	"github.com/Tarasusrus/url-shortener/internal/app/models"
	"github.com/Tarasusrus/url-shortener/internal/app/stores"
	"github.com/Tarasusrus/url-shortener/internal/logger"
	"github.com/gin-gonic/gin"
	"github.com/mailru/easyjson"
	"go.uber.org/zap"
	"mime"
	"net/http"
)

const ExpectedContentType = "application/json"

func HandleJSONPost(writer gin.ResponseWriter, request *http.Request, store *stores.Store, config *configs.FlagConfig) {
	// Проверка ContentType на соответствие text/plain.
	mediaType, _, err := mime.ParseMediaType(request.Header.Get("Content-Type"))

	//todo вынести в функцию повторяющийся код.
	if err != nil || mediaType != ExpectedContentType {
		logger.Log.Debug("Invalid content type",
			zap.Error(err),
			zap.String("expected", ExpectedContentType),
			zap.String("got", mediaType)) // Логирование ожидаемого и полученного типов контента
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	var req models.ShortenRequest
	if err := easyjson.UnmarshalFromReader(request.Body, &req); err != nil {
		logger.Log.Debug("error in UnmarshalFromReader",
			zap.Error(err))
		http.Error(writer, err.Error(), http.StatusBadRequest)
	}

	defer func() {
		if err := request.Body.Close(); err != nil {
			logger.Log.Error("Failed to close request body", zap.Error(err))
		}
	}()

	shortURLId := store.Set(req.URL)
	logger.Log.Info("Short URL created", zap.String("shortURLId", shortURLId))

	resp := models.ShortenResponse{ShortURL: shortURLId}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	if _, err := easyjson.MarshalToWriter(resp, writer); err != nil {
		logger.Log.Error("Failed to marshal response", zap.Error(err))
	}
}
