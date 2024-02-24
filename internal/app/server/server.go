package server

import (
	"fmt"
	"github.com/Tarasusrus/url-shortener/internal/app/configs"
	"github.com/Tarasusrus/url-shortener/internal/app/handlers"
	"github.com/Tarasusrus/url-shortener/internal/app/stores"
	"github.com/Tarasusrus/url-shortener/internal/logger"
	"github.com/gin-gonic/gin"
)

func Run() error {
	config, err := configs.NewFlagConfig()
	if err != nil {
		return fmt.Errorf("failed to load flag config: %w", err)
	}
	if err := logger.Initialize(config.LogLevel); err != nil {
		return fmt.Errorf("failed to initialize logger: %w", err)
	}
	store := stores.NewStore()
	r := gin.Default()

	// Регистрация middleware
	r.Use(logger.RequestLoggerMiddleware())

	r.GET("/:id", func(c *gin.Context) {
		handlers.HandleGet(c.Writer, c.Request, store)
	})
	r.POST("/", func(c *gin.Context) {
		handlers.HandlePost(c.Writer, c.Request, store, config)
	})

	if err := r.Run(config.GetAddress()); err != nil {
		return fmt.Errorf("failed to run server: %w", err)
	}
	return nil
}
