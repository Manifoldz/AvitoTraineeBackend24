package banner

type Banner struct {
	Feature_id int     `json:"feature_id" binding:"required"`
	Tag_ids    []int   `json:"tag_ids" binding:"required"`
	Content    Content `json:"content" binding:"required"`
	Is_active  bool    `json:"is_active" binding:"required"`
}
