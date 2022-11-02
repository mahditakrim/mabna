package rest

import (
	"github.com/kataras/iris/v12"
	"github.com/mahditakrim/mabna/service"
)

func getLatestTradesHandler(trade service.Trade) func(iris.Context) {
	return func(ctx iris.Context) {

		trades, err := trade.GetLatestTradeOfIstruments(ctx.Request().Context())
		if err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(struct{ Err string }{err.Error()})
			return
		}

		ctx.JSON(trades)
	}
}
