package handler

import "github.com/gin-gonic/gin"

type Handler struct {
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
