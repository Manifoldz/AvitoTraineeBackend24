package banner

type Content struct {
	Title string `json:"title" binding:"required"`
	Text  string `json:"text" binding:"required"`
	URL   string `json:"url" binding:"required"`
}
