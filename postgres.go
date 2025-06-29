package ec2test

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

func NewPostgresConn(postgreURI string) *pgxpool.Pool {
	pool, err := pgxpool.New(context.Background(), postgreURI)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create pool of connections")
	}

	return pool
}
