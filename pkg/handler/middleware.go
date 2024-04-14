package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) userIndentity(c *gin.Context) {
	header := c.GetHeader("token")

	// Проверка токена на пустоту
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "Пользователь не авторизован")
		return
	}

	// Валидность токена и определение роли пользователя
	_, err := h.services.AuthenticateToken(header)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "Пользователь не авторизован")
		return
	}

	c.Next() //переход к след эндпоинту
}

func (h *Handler) adminIndentity(c *gin.Context) {
	header := c.GetHeader("token")

	// Проверка токена на пустоту
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "Пользователь не авторизован")
		return
	}

	// Валидность токена и определение роли пользователя
	role, err := h.services.AuthenticateToken(header)
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "Пользователь не авторизован")
		return
	}
	// Если роль не админ, запретить доступ
	if role != "admin" {
		newErrorResponse(c, http.StatusForbidden, "Пользователь не имеет доступа")
		return
	}

	c.Next() //переход к след эндпоинту
}
