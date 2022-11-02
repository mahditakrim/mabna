package rest

import (
	"context"
	"errors"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/mahditakrim/mabna/service"
	"github.com/mahditakrim/mabna/transport"
)

type rest struct {
	router *iris.Application
}

func NewHttp(service service.Trade) (transport.Transport, error) {

	if service == nil {
		return nil, errors.New("nil service reference")
	}

	router := iris.New()
	router.PartyFunc("/mabna", func(mabna iris.Party) {
		mabna.Get("/latest_trades", getLatestTradesHandler(service))
	})

	return &rest{router}, nil
}

func (http *rest) Run(addr string) error {

	return http.router.Listen(addr, iris.WithoutServerError(iris.ErrServerClosed))
}

func (http *rest) Shutdown() error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return http.router.Shutdown(ctx)
}
