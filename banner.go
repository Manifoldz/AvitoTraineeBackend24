package banner

import "github.com/lib/pq"

type Banner struct {
	BannerId   int `json:"banner_id" db:"banner_id"`
	Feature_id int `json:"feature_id" db:"feature_id" binding:"required"`
	// Tag_ids пришлось поменять с обычного массива интов т.к. не дает прочитать данные из Postgres
	Tag_ids pq.Int64Array `json:"tag_ids" db:"tag_ids" binding:"required"`
	Content Content       `json:"content" db:"content" binding:"required"`
	// Is_active обязательно указатель надо, иначе не биндится - есть на github обсуждение
	Is_active  *bool  `json:"is_active" db:"is_active" binding:"required"`
	Created_at string `json:"created_at" db:"created_at"`
	Updated_at string `json:"updated_at" db:"updated_at"`
}

type QueryParams struct {
	FeatureID *int
	TagID     *int
	Limit     *int
	Offset    *int
}
