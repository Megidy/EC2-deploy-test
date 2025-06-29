package main

import (
	"os"
	"os/signal"
	"syscall"

	ec2test "github.com/Megidy/ec2-test"
	"github.com/rs/zerolog/log"
)

func main() {

	cfg := ec2test.NewConifg()

	dbPool := ec2test.NewPostgresConn(cfg.PostgresURI)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	server := ec2test.NewServer(cfg.HttpServerPort, dbPool, cfg.ApiKey)

	err := server.Run()
	if err != nil {
		return
	}

	go func() {
		<-sigs
		log.Info().Msg("Received signal to shutdown")

		err = server.Shutdown()
		if err != nil {
			log.Err(err).Msg("Failed to shutdown application")
		}

		err := server.Shutdown()
		if err != nil {
			log.Error().Err(err).Msg("failed to shutdown")
		}
	}()

}
