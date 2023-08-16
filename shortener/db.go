package shortener

import (
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func SetPool(newPool *pgxpool.Pool) error {

	if newPool == nil {
		return errors.New("cannot assign nil pool")
	}

	Pool = newPool

	return nil
}
