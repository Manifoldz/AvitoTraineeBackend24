package banner

type Banner struct {
	Banner_id  int   `json:"banner_id"`
	Feature_id int   `json:"feature_id"`
	Tag_ids    []int `json:"tag_ids"`
	Content    `json:"content"`
	Is_active  bool   `json:"is_active"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}
