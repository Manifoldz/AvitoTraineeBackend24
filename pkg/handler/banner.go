package handler

import (
	"net/http"

	banner "github.com/Manifoldz/AvitoTraineeBackend24"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getAllBannersByFeatureAndOrTag(c *gin.Context) {

}

func (h *Handler) createBanner(c *gin.Context) {
	var input banner.Banner

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Некорректные данные")
		return
	}

	id, err := h.services.Banner.Create(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, map[string]interface{}{
		"banner_id": id,
	})
}

func (h *Handler) updateContentBannerById(c *gin.Context) {

}

func (h *Handler) deleteBannerById(c *gin.Context) {

}
