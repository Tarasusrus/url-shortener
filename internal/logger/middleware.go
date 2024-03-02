package logger

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
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

func GzipDecodeMiddleware(c *gin.Context) {
	if c.GetHeader("Content-Encoding") == "gzip" {
		reader, err := gzip.NewReader(c.Request.Body)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		c.Request.Body = reader
		c.Next()
	} else {
		c.Next()
	}
}

func GzipEncodeMiddleware(c *gin.Context) {
	if strings.Contains(c.GetHeader("Accept-Encoding"), "gzip") {
		// Создаем gzip writer
		gz := gzip.NewWriter(c.Writer)
		defer gz.Close()

		// Заменяем writer контекста на gzip writer
		c.Writer = &gzipResponseWriter{Writer: gz, ResponseWriter: c.Writer}

		// Установка заголовка Content-Encoding
		c.Header("Content-Encoding", "gzip")
	}
	c.Next()
}

type gzipResponseWriter struct {
	gin.ResponseWriter
	Writer io.Writer
}

// Переопределение метода Write, чтобы использовать gzip.Writer
func (g *gzipResponseWriter) Write(data []byte) (int, error) {
	return g.Writer.Write(data)
}
