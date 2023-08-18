package main

import (
	"capt-shortener/shortener"
	"capt-shortener/util"
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

func main() {
	dbUrl := fmt.Sprintf("host=%s dbname=%s user=%s password=%s port=%s", util.Conf.DBHost, util.Conf.DBName, util.Conf.DBUser, util.Conf.DBPassword, util.Conf.DBPort)
	pool, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		log.Error().Err(err).Msg("Unable to connect to db")
	}
	shortener.SetPool(pool)

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", util.Conf.RedisHost, util.Conf.RedisPort),
		Password: util.Conf.RedisPassword,
		DB:       0,
	})
	shortener.SetRdbClient(rdb)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Mount("/", shortener.Router())
	serverStr := fmt.Sprintf(":%s", util.Conf.ServerPort)
	log.Info().Msg("Starting up server " + serverStr)
	if err := http.ListenAndServe(serverStr, r); err != nil {
		log.Fatal().Err(err).Msg("Failed to start the server")
		return
	}
}
