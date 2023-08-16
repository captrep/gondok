package shortener

import (
	"encoding/json"
	"time"
)

func (l Link) MarshalJSON() ([]byte, error) {
	var j struct {
		ID        string    `json:"id"`
		ShortURL  string    `json:"short_url"`
		LongURL   string    `json:"long_url"`
		CreatedAt time.Time `json:"created_at"`
	}
	j.ID = l.ID
	j.ShortURL = l.ShortURL
	j.LongURL = l.LongURL
	j.CreatedAt = l.CreatedAt

	return json.Marshal(j)
}
