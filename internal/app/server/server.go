package server

import (
	"github.com/Tarasusrus/url-shortener/helpers"
	"github.com/Tarasusrus/url-shortener/internal/app/configs"
	"github.com/Tarasusrus/url-shortener/internal/app/handlers"
	"github.com/Tarasusrus/url-shortener/internal/app/stores"
	"github.com/Tarasusrus/url-shortener/internal/logger"
	"github.com/gin-gonic/gin"
	"log"
)

func Run() {
	config, err := configs.NewFlagConfig()
	if err != nil {
		helpers.LogError(err)
	}
	err = logger.Initialize(config.LogLevel)
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
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

	log.Fatal(r.Run(config.GetAddress()))
}
