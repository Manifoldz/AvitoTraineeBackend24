package banner

import (
	"encoding/json"
	"errors"
)

type Content struct {
	Title string `json:"title" binding:"required"`
	Text  string `json:"text" binding:"required"`
	URL   string `json:"url" binding:"required"`
}

// Scan реализует интерфейс sql.Scanner для Content, чтобы
// можно было прочитать JSONB данные и преобразовать их в структуру Content.
func (c *Content) Scan(src interface{}) error {
	if src == nil {
		return nil
	}

	switch src := src.(type) {
	case []byte:
		return json.Unmarshal(src, c)
	default:
		return errors.New("incompatible type for Content")
	}
}
