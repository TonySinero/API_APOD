package handler

import (
	_ "github.com/apod/docs"
	"github.com/apod/internal/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Use(
		h.CorsMiddleware,
	)

	album := router.Group("/album")
	{
		album.POST("/", h.createAlbum)
		album.GET("/images", h.getAlbumFromDB)
		album.GET("/filter", h.getWithFilter)
	}

	return router
}
