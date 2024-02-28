// Package server содержит основную логику для запуска HTTP-сервера приложения.
// Он использует пакет Gin для роутинга и обработки HTTP-запросов.
package server

import (
	"fmt"

	"github.com/Tarasusrus/url-shortener/internal/app/configs"
	"github.com/Tarasusrus/url-shortener/internal/app/handlers"
	"github.com/Tarasusrus/url-shortener/internal/app/stores"
	"github.com/Tarasusrus/url-shortener/internal/logger"
	"github.com/gin-gonic/gin"
)

// Run инициализирует и запускает HTTP-сервер приложения.
// Эта функция сначала загружает конфигурацию сервера, используя флаги командной строки и переменные окружения.
// Затем она инициализирует логгер, создает хранилище для URL и настраивает роутинг с помощью пакета Gin.
// Все входящие запросы обрабатываются с использованием зарегистрированных обработчиков для GET и POST запросов.
// В случае возникновения ошибок на любом этапе инициализации, функция вернет соответствующую ошибку.
// При успешной инициализации сервер запускается и остается в работе, ожидая входящие HTTP-запросы.
func Run() error {
	config, err := configs.NewFlagConfig()
	if err != nil {
		return fmt.Errorf("failed to load flag config: %w", err)
	}

	if err := logger.Initialize(config.LogLevel); err != nil {
		return fmt.Errorf("failed to initialize logger: %w", err)
	}

	store := stores.NewStore()
	router := gin.Default()

	// Регистрация middleware
	router.Use(logger.RequestLoggerMiddleware())

	router.GET("/:id", func(c *gin.Context) {
		handlers.HandleGet(c.Writer, c.Request, store)
	})
	router.POST("/", func(c *gin.Context) {
		handlers.HandlePost(c.Writer, c.Request, store, config)
	})
	router.POST("/api/shorten", func(c *gin.Context) {
		handlers.HandleJSONPost(c.Writer, c.Request, store, config)
	})

	if err := router.Run(config.GetAddress()); err != nil {
		return fmt.Errorf("failed to run server: %w", err)
	}
	return nil
}
