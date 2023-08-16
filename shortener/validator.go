package shortener

import (
	"errors"
	"net/url"
)

func Validate(shortURL, longURL string) error {
	if len(shortURL) == 0 {
		return errors.New("shorturl: is empty")
	}
	if len(longURL) == 0 {
		return errors.New("longurl: is empty")
	}

	_, err := url.ParseRequestURI(longURL)
	if err != nil {
		return errors.New("not a valid URL")
	}
	return nil
}
