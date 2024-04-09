package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	user_banner := router.Group("/user_banner")
	{
		user_banner.GET("/")
	}

	banner := router.Group("/banner")
	{
		banner.GET("/")
		banner.POST("/")
		banner.PATCH("/:id")
		banner.DELETE("/:id")
	}

	return router
}
