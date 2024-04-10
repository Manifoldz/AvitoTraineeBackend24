package repository

import "github.com/jmoiron/sqlx"

type Banner interface {
}

type UserBanner interface {
}

type Repository struct {
	Banner
	UserBanner
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
