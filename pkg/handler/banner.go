package handler

import (
	"net/http"
	"strconv"

	banner "github.com/Manifoldz/AvitoTraineeBackend24"
	"github.com/gin-gonic/gin"
)

type getAllBannersByFeatureAndOrTagResponse struct {
	Data []banner.Banner `json:"data"`
}

func (h *Handler) getAllBannersByFeatureAndOrTag(c *gin.Context) {
	// Получение параметров фильтрации и пагинации из запроса

	//использую указатели, чтобы не запрещать использовать отрицательные и нулевые значения
	var newQueryParam banner.QueryParams
	var err error

	if newQueryParam.FeatureID, err = getParametr(c, "feature_id"); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Некорректные данные")
		return
	}
	if newQueryParam.TagID, err = getParametr(c, "tag_id"); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Некорректные данные")
		return
	}
	if newQueryParam.Limit, err = getParametr(c, "limit"); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Некорректные данные")
		return
	}
	if newQueryParam.Offset, err = getParametr(c, "offset"); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Некорректные данные")
		return
	}
	banners, err := h.services.Banner.GetAllFiltered(&newQueryParam)
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

// Функция для извлечения параметра из запроса
func getParametr(c *gin.Context, paramName string) (*int, error) {
	paramValue := c.Query(paramName)
	if paramValue == "" {
		return nil, nil // Параметр не был передан
	}
	intValue, err := strconv.Atoi(paramValue)
	if err != nil {
		return nil, err // Параметр не является допустимым целым числом
	}
	return &intValue, nil
}
