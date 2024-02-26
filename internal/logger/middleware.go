package logger

import (
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// RequestLoggerMiddleware создает промежуточное ПО (middleware) для логирования запросов.
// Оно логирует метод, путь, статус ответа, размер ответа и задержку обработки каждого запроса.
func RequestLoggerMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		startTime := time.Now()

		// Обработка запроса
		context.Next()

		// Вычисляем продолжительность обработки запроса
		latency := time.Since(startTime)
		statusCode := context.Writer.Status()
		responseSize := context.Writer.Size()

		// Логируем информацию о запросе и ответе
		Log.Info("request",
			zap.String("method", context.Request.Method),
			zap.String("path", context.Request.URL.Path),
			zap.Int("status", statusCode),
			zap.Int64("response_size", int64(responseSize)),
			zap.Duration("latency", latency),
		)
	}
}
