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

	// //все что ниже до упоминания удалить не забыть!!!
	// bodyBytes, err := io.ReadAll(c.Request.Body)
	// if err != nil {
	// 	// Обработка ошибки чтения тела
	// 	newErrorResponse(c, http.StatusBadRequest, "Ошибка чтения тела запроса")
	// 	return
	// }

	// // Логирование тела запроса
	// bodyString := string(bodyBytes)
	// fmt.Println("Тело запроса:", bodyString)

	// //до этого места удаляем

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error()) // заменить на "Некорректные данные")
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
