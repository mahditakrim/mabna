package service

import (
	"context"
	"errors"

	"github.com/mahditakrim/mabna/entity"
	"github.com/mahditakrim/mabna/repository"
)

type Trade interface {
	Validate(*entity.Trade) error
	GetLatestTradeOfIstruments(context.Context) ([]entity.Trade, error)
}

type trade struct {
	repo repository.Repository
}

func NewTradeService(repo repository.Repository) (Trade, error) {

	if repo == nil {
		return nil, errors.New("nil respository reference")
	}

	return &trade{repo}, nil
}

func (*trade) Validate(book *entity.Trade) error {

	//NOTE: would use in case of adding new trades
	return nil
}

func (l *trade) GetLatestTradeOfIstruments(ctx context.Context) ([]entity.Trade, error) {

	return l.repo.GetLatestTrades(ctx)
}
