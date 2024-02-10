package server

import (
	"github.com/Tarasusrus/url-shortener/internal/app/configs"
	"github.com/Tarasusrus/url-shortener/internal/app/handlers"
	"github.com/Tarasusrus/url-shortener/internal/app/stores"
	"github.com/gin-gonic/gin"
	"log"
)

func Run() {
	store := stores.NewStore()
	config := configs.NewFlagConfig()

	r := gin.Default()

	r.GET("/:id", func(c *gin.Context) {
		handlers.HandleGet(c.Writer, c.Request, store)
	})
	r.POST("/", func(c *gin.Context) {
		handlers.HandlePost(c.Writer, c.Request, store, config)
	})
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	switch r.Method {
	//	case http.MethodGet:
	//		handlers.HandleGet(w, r, store)
	//	case http.MethodPost:
	//		handlers.HandlePost(w, r, store)
	//	default:
	//		w.WriteHeader(http.StatusBadRequest) // На любой некорректный запрос сервер должен возвращать ответ с кодом 400
	//	}
	//})
	log.Fatal(r.Run(config.GetAddress()))
}
