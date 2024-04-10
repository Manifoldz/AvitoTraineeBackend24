package service

import (
	"errors"
	"strings"
)

// AuthenticateToken проверяет токен и возвращает роль пользователя.
func (s *Service) AuthenticateToken(token string) (string, error) {

	if strings.Contains(token, "admin") {
		return "admin", nil
	} else if strings.Contains(token, "user") {
		return "user", nil
	}

	return "", errors.New("wrong token")
}
