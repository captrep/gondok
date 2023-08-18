package shortener

import (
	"errors"

	"github.com/redis/go-redis/v9"
)

var Rdb *redis.Client

func SetRdbClient(newRdbClient *redis.Client) error {
	if newRdbClient == nil {
		return errors.New("nil client")
	}
	Rdb = newRdbClient
	return nil
}
