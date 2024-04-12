package service

import (
	banner "github.com/Manifoldz/AvitoTraineeBackend24"
	"github.com/Manifoldz/AvitoTraineeBackend24/pkg/repository"
)

type Banner interface {
	Create(ban banner.Banner) error
}

type UserBanner interface {
}

type Service struct {
	Banner
	UserBanner
}

func NewService(*repository.Repository) *Service {
	return &Service{}
}
