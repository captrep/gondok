package shortener

import "time"

type Link struct {
	ID        string
	ShortURL  string
	LongURL   string
	CreatedAt time.Time
}
