package service

import (
	banner "github.com/Manifoldz/AvitoTraineeBackend24"
	"github.com/Manifoldz/AvitoTraineeBackend24/pkg/repository"
)

type Banner interface {
	Create(ban banner.Banner) (int, error)
	GetAllFiltered(requestParam *banner.RequestParams) ([]banner.Banner, error)
}

type UserBanner interface {
}

type Service struct {
	Banner
	UserBanner
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Banner: NewBannerBannerService(repos.Banner),
	}
}
