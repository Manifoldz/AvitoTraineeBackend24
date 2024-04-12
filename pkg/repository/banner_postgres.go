package repository

import (
	"fmt"

	banner "github.com/Manifoldz/AvitoTraineeBackend24"
	"github.com/jmoiron/sqlx"
)

type BannerPostgres struct {
	db *sqlx.DB
}

func NewBannerPostgres(db *sqlx.DB) *BannerPostgres {
	return &BannerPostgres{db: db}
}

func (r *BannerPostgres) Create(banner.Banner) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var id int
	createBannerQuery := fmt.Sprintf("INSERT INTO %s (content, is_active,) VALUES ($1, $2) RETURNING id", bannersTable)
	row := tx.QueryRow(createBannerQuery, banner.Content, banner.Is_active)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	createTagQuery := fmt.Sprintf("INSERT INTO %s (id) VALUES ($1)", tagsTable)

	createFeatureQuery := fmt.Sprintf("INSERT INTO %s (id) VALUES ($1)", featuresTable)
}
