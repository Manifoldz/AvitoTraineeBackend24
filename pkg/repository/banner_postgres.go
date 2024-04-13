package repository

import (
	"database/sql"
	"encoding/json"
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

func (r *BannerPostgres) Create(ban banner.Banner) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	//проверим существует ли уже фича, если нет то добавим
	featureExists, err := r.ensureExists(tx, featuresTable, ban.Feature_id)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	if !featureExists {
		createFeatureQuery := fmt.Sprintf("INSERT INTO %s (id) VALUES ($1)", featuresTable)
		_, err = tx.Exec(createFeatureQuery, ban.Feature_id)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	//проверим теги и аналогично добавляем если их нет, но при этом фича должна была отсутствовать!
	//иначе получается что данный баннер в базе уже есть и нужно было запрос на обновление посылать!
	for _, val := range ban.Tag_ids {
		tagExists, err := r.ensureExists(tx, tagsTable, val)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
		if !tagExists {
			createTagQuery := fmt.Sprintf("INSERT INTO %s (id) VALUES ($1)", tagsTable)
			_, err := tx.Exec(createTagQuery, val)
			if err != nil {
				tx.Rollback()
				return 0, err
			}
		} else if featureExists {
			tx.Rollback()
			return 0, fmt.Errorf("error occured while creating new banner - banner already exists")
		}
	}

	contentJSON, err := json.Marshal(ban.Content)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	var id int
	createBannerQuery := fmt.Sprintf("INSERT INTO %s (content, is_active) VALUES ($1, $2) RETURNING id", bannersTable)
	row := tx.QueryRow(createBannerQuery, contentJSON, ban.Is_active)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createBannerFeatureTagQuery := fmt.Sprintf("INSERT INTO %s (banner_id, feature_id, tag_id) VALUES ($1, $2, $3)", bannerFeatureTagTable)
	for _, val := range ban.Tag_ids {
		_, err = tx.Exec(createBannerFeatureTagQuery, id, ban.Feature_id, val)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	return id, tx.Commit()
}

// метод проверяет - а не существует ли уже такой id?
func (r *BannerPostgres) ensureExists(tx *sql.Tx, table string, id int) (bool, error) {
	var exists bool
	checkExistsQuery := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE id = $1)", table)
	row := tx.QueryRow(checkExistsQuery, id)
	err := row.Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
