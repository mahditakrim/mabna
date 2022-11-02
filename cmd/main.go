package main

import (
	"os"

	"github.com/mahditakrim/mabna/cmd/config"
	"github.com/mahditakrim/mabna/internal/shutdown"
	"github.com/mahditakrim/mabna/repository"
	"github.com/mahditakrim/mabna/service"
	"github.com/mahditakrim/mabna/transport/rest"
)

func main() {

	// init config
	err := config.Init()
	if err != nil {
		panic(err)
	}

	// make dependencies in order
	repository, err := repository.NewPostgres(
		config.Get().DB.Postgres.Addr,
		config.Get().DB.Postgres.Username,
		config.Get().DB.Postgres.Password,
		config.Get().DB.Postgres.Db,
	)
	if err != nil {
		panic(err)
	}

	tradeService, err := service.NewTradeService(repository)
	if err != nil {
		panic(err)
	}

	httpServer, err := rest.NewHttp(tradeService)
	if err != nil {
		panic(err)
	}

	// run app
	go func() {
		err := httpServer.Run(config.Get().Transport.HttpAddr)
		if err != nil {
			panic(err)
		}
	}()

	// gracefull shutdown
	err = shutdown.Graceful(
		func() error {
			err := httpServer.Shutdown()
			if err != nil {
				return err
			}
			return repository.Close()
		},
		make(chan os.Signal, 1),
	)
	if err != nil {
		panic(err)
	}
}
