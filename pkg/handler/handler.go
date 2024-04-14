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

	user_banner := router.Group("/user_banner", h.userIndentity)
	{
		user_banner.GET("/", h.getUserBanner)
	}

	admin_banner := router.Group("/banner", h.adminIndentity)
	{
		admin_banner.GET("/", h.getAllBannersByFeatureAndOrTag)
		admin_banner.POST("/", h.createBanner)
		admin_banner.PATCH("/:id", h.updateContentBannerById)
		admin_banner.DELETE("/:id", h.deleteBannerById)
	}

	return router
}
