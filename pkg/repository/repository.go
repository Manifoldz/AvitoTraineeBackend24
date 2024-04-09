package repository

type Banner interface {
}

type UserBanner interface {
}

type Repository struct {
	Banner
	UserBanner
}

func NewRepository() *Repository {
	return &Repository{}
}
