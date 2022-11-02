package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/mahditakrim/mabna/entity"
)

type Repository interface {
	Close() error
	GetLatestTrades(context.Context) ([]entity.Trade, error)
}

type postgres struct {
	conn *sql.DB
}

func NewPostgres(addr, username, password, database string) (Repository, error) {

	conn, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		username,
		password,
		addr,
		database))
	if err != nil {
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	return &postgres{conn}, nil
}

func (db *postgres) GetLatestTrades(ctx context.Context) ([]entity.Trade, error) {

	rows, err := db.conn.QueryContext(ctx, `SELECT DISTINCT ON (instruments.name) instruments.name, trades.* FROM trades 
	INNER JOIN instruments ON trades.instrument_id = instruments.id
	ORDER BY instruments.name, trades.dateEn DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var trades []entity.Trade
	for rows.Next() {
		var trade entity.Trade
		err = rows.Scan(&trade.InstrumentName, &trade.ID, &trade.InstrumentID, &trade.Date, &trade.Open, &trade.High,
			&trade.Low, &trade.Close)
		if err != nil {
			return nil, err
		}
		trades = append(trades, trade)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	if trades == nil {
		return nil, errors.New("no trades")
	}

	return trades, nil
}

func (db *postgres) Close() error {

	return db.conn.Close()
}
