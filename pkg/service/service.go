package service

import "github.com/Manifoldz/AvitoTraineeBackend24/pkg/repository"

type Banner interface {
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
