package banner

type Banner struct {
	BannerId   int     `json:"banner_id"`
	Feature_id int     `json:"feature_id" binding:"required"`
	Tag_ids    []int   `json:"tag_ids" binding:"required"`
	Content    Content `json:"content" binding:"required"`
	Is_active  bool    `json:"is_active" binding:"required"`
	Created_at string  `json:"created_at"`
	Updated_at string  `json:"updated_at"`
}

type QueryParams struct {
	FeatureID *int
	TagID     *int
	Limit     *int
	Offset    *int
}
