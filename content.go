package banner

type Content struct {
	Id    int    `json:"-"`
	Title string `json:"title"`
	Text  string `json:"text"`
	URL   string `json:"url"`
}
