package shortener

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

func CreateLink(ctx context.Context, shortURL, longURL string) (Link, error) {
	link, err := NewLink(shortURL, longURL)
	if err != nil {
		return link, err
	}

	tx, err := Pool.Begin(ctx)
	if err != nil {
		return link, err
	}

	err = Save(ctx, tx, link)
	if err != nil {
		tx.Rollback(ctx)
		return link, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return link, err
	}

	err = SetURL(ctx, shortURL, longURL)
	if err != nil {
		log.Error().Err(err).Msg("error set url redis from get url service")
	}
	return link, nil
}

func GetLink(ctx context.Context, shortURL string) (Link, error) {
	var link Link

	longURL, err := GetURL(ctx, shortURL)
	if err != nil {
		if err != redis.Nil && longURL != "" {
			link.LongURL = longURL
			return link, nil
		}
		log.Debug().Err(err).Msg("fail getting url from redis")
	}

	tx, err := Pool.Begin(ctx)
	if err != nil {
		log.Debug().Err(err).Msg("error when begin tx")
		return link, err
	}

	link, err = Find(ctx, tx, shortURL)
	if err != nil {
		log.Debug().Err(err).Msg("err")
		tx.Rollback(ctx)
		return link, err
	}

	err = SetURL(ctx, shortURL, link.LongURL)
	if err != nil {
		log.Error().Err(err).Msg("error set url redis from get url service")
		err = nil
	}

	err = tx.Commit(ctx)
	if err != nil {
		return link, err
	}

	return link, nil
}
