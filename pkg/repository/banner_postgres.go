package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"

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
	featureExists, err := r.ensureExists(tx, featuresTable, int64(ban.Feature_id))
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

// метод проверяет - а не существует ли уже такой id в таблице?
func (r *BannerPostgres) ensureExists(tx *sql.Tx, table string, id int64) (bool, error) {
	var exists bool
	checkExistsQuery := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE id = $1)", table)
	row := tx.QueryRow(checkExistsQuery, id)
	err := row.Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (r *BannerPostgres) GetAllFiltered(requestParam *banner.RequestParams) ([]banner.Banner, error) {
	var queryBuilder strings.Builder
	var args []interface{}
	var argCounter int = 1

	// Добавление выбора, исх.таблиц, объединения
	queryBuilder.WriteString(`
	SELECT b.id AS banner_id, bft.feature_id, array_agg(bft.tag_id) AS tag_ids, 
		b.content, b.is_active, b.created_at, b.updated_at`)
	queryBuilder.WriteString(fmt.Sprintf(" FROM %s b JOIN %s bft ON b.id = bft.banner_id", bannersTable, bannerFeatureTagTable))

	// Добавление подзапроса с фильтрацией и фильтрации во внешнем запросе
	queryBuilder.WriteString(fmt.Sprintf(" WHERE b.id IN (SELECT banner_id FROM %s", bannerFeatureTagTable))

	// Добавление фильтрации по feature_id, если он передан
	if requestParam.FeatureID != nil {
		queryBuilder.WriteString(fmt.Sprintf(" WHERE feature_id = $%d", argCounter))
		args = append(args, *requestParam.FeatureID)
		argCounter++
	}

	// Добавление фильтрации по tag_id, если он передан и уже есть условие WHERE
	if requestParam.TagID != nil {
		if requestParam.FeatureID != nil {
			queryBuilder.WriteString(" AND")
		} else {
			queryBuilder.WriteString(" WHERE")
		}
		queryBuilder.WriteString(fmt.Sprintf(" tag_id = $%d", argCounter))
		args = append(args, *requestParam.TagID)
		argCounter++
	}

	// Добавление группировки
	queryBuilder.WriteString(`) 
	GROUP BY
		b.id,
		bft.feature_id,
		b.content,
		b.is_active,
		b.created_at,
		b.updated_at`)

	// Добавление пагинации
	insertLimit := 10 // по умолчанию лимит
	insertOffset := 0 // по умолчанию оффсет
	if requestParam.Limit != nil {
		insertLimit = *requestParam.Limit
	}
	if requestParam.Offset != nil {
		insertOffset = *requestParam.Offset
	}
	queryBuilder.WriteString(fmt.Sprintf(" LIMIT $%d OFFSET $%d", argCounter, argCounter+1))
	args = append(args, insertLimit, insertOffset)

	// Выполнение запроса
	finalQuery := queryBuilder.String()
	var banners []banner.Banner
	err := r.db.Select(&banners, finalQuery, args...)

	return banners, err
}
