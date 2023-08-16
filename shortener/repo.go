package shortener

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
)

func Save(ctx context.Context, tx pgx.Tx, link Link) error {
	q := `INSERT INTO links (id, short_url, long_url, created_at) VALUES ($1, $2, $3, $4)`
	_, err := tx.Exec(ctx, q, link.ID, link.ShortURL, link.LongURL, link.CreatedAt)
	if err != nil {
		log.Err(err).Msg("error save to db")
		return err
	}

	return nil
}

func Find(ctx context.Context, tx pgx.Tx, shortURL string) (Link, error) {
	q := `SELECT id, short_url, long_url, created_at FROM links WHERE short_url = $1`
	row := tx.QueryRow(ctx, q, shortURL)
	var link Link
	if err := row.Scan(&link.ID, &link.ShortURL, &link.LongURL, &link.CreatedAt); err != nil {
		return link, err
	}
	return link, nil
}
