package shortener

import (
	"time"

	"github.com/google/uuid"
)

type Link struct {
	ID        string
	ShortURL  string
	LongURL   string
	CreatedAt time.Time
}

func NewLink(shortURL, longURL string) (Link, error) {
	if err := Validate(shortURL, longURL); err != nil {
		return Link{}, err
	}

	link := Link{
		ID:        uuid.NewString(),
		ShortURL:  shortURL,
		LongURL:   longURL,
		CreatedAt: time.Now(),
	}

	return link, nil
}
