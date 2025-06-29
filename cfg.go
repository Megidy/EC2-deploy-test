package ec2test

import (
	"github.com/caarlos0/env"
	"github.com/rs/zerolog/log"
)

type Config struct {
	HttpServerPort string `env:"HTTP_SERVER_PORT,required"`
	PostgresURI    string `env:"POSRGRES_URI,required"`
	ApiKey         string `env:"API_KEY,required"`
}

func NewConifg() *Config {
	var cfg Config

	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to parse config")
	}
	return &cfg
}
