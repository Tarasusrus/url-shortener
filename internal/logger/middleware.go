// logger/middleware.go

package logger

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func RequestLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		// Обработка запроса
		c.Next()

		// Вычисляем продолжительность обработки запроса
		latency := time.Now().Sub(startTime)
		statusCode := c.Writer.Status()
		responseSize := c.Writer.Size()

		// Логируем информацию о запросе и ответе
		Log.Info("request",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", statusCode),
			zap.Int64("response_size", int64(responseSize)),
			zap.Duration("latency", latency),
		)
	}
}
