package handler

import (
	"github.com/Manifoldz/AvitoTraineeBackend24/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	user_banner := router.Group("/user_banner")
	{
		user_banner.GET("/", h.getUserBanner)
	}

	banner := router.Group("/banner")
	{
		banner.GET("/", h.getAllBannersByFeatureAndOrTag)
		banner.POST("/", h.createBanner)
		banner.PATCH("/:id", h.updateContentBannerById)
		banner.DELETE("/:id", h.deleteBannerById)
	}

	return router
}
