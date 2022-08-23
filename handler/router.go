package handler

import (
	"github.com/apod/service"
	"github.com/gin-gonic/gin"
)

// Handler type replies for handling gin server requests.
type Handler struct {
	services *service.Service
}

// NewHandler function create handler.
func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(
		h.CorsMiddleware,
	)

	album := router.Group("/album")
	{
		album.POST("/", h.createAlbum)
		album.GET("/dt", h.getByDate)
		album.GET("/images", h.getAlbumFromDB)
		album.GET("/filter", h.getWithFilter)
	}

	return router
}
