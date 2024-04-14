package handler

import (
	"net/http"

	banner "github.com/Manifoldz/AvitoTraineeBackend24"
	"github.com/gin-gonic/gin"
)

type getAllBannersByFeatureAndOrTagResponse struct {
	Data []banner.Banner `json:"data"`
}

func (h *Handler) getAllBannersByFeatureAndOrTag(c *gin.Context) {
	// Получение параметров фильтрации и пагинации из запроса

	//использую указатели, чтобы не запрещать использовать отрицательные и нулевые значения
	var newRequestParam banner.RequestParams
	var err error

	if err := c.ShouldBindJSON(&newRequestParam); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Некорректные данные")
		return
	}

	banners, err := h.services.Banner.GetAllFiltered(&newRequestParam)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllBannersByFeatureAndOrTagResponse{
		Data: banners,
	})
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
