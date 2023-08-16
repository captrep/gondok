package main

import (
	"capt-shortener/shortener"
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

func main() {
	pool, err := pgxpool.New(context.Background(), "host=localhost dbname=shortener user=postgres password=root port=5431")
	if err != nil {
		log.Error().Err(err).Msg("Unable to connect to db")
	}
	shortener.SetPool(pool)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "application/json")
		w.WriteHeader(200)
		resp := map[string]interface{}{
			"code": 200,
			"msg":  "Ok",
		}
		json.NewEncoder(w).Encode(resp)
	})
	// r.Mount("/api/v1", api.Router())
	log.Info().Msg("Starting up server 127.0.0.1:8000")

	if err := http.ListenAndServe("127.0.0.1:8000", r); err != nil {
		log.Fatal().Err(err).Msg("Failed to start the server")
		return
	}
}
