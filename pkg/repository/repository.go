package repository

import (
	banner "github.com/Manifoldz/AvitoTraineeBackend24"
	"github.com/jmoiron/sqlx"
)

type Banner interface {
	Create(ban banner.Banner) (int, error)
	GetAllFiltered(newRequestParam *banner.RequestParams) ([]banner.Banner, error)
}

type UserBanner interface {
}

type Repository struct {
	Banner
	UserBanner
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Banner: NewBannerPostgres(db),
	}
}
