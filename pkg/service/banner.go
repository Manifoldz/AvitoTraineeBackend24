package service

import (
	banner "github.com/Manifoldz/AvitoTraineeBackend24"
	"github.com/Manifoldz/AvitoTraineeBackend24/pkg/repository"
)

type BannerBannerService struct {
	repo repository.Banner
}

func NewBannerBannerService(repo repository.Banner) *BannerBannerService {
	return &BannerBannerService{repo: repo}
}

func (s *BannerBannerService) Create(ban banner.Banner) (int, error) {
	return s.repo.Create(ban)
}
