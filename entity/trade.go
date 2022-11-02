package entity

import "time"

type Trade struct {
	InstrumentName string    `json:"instrument_name"`
	ID             int64     `json:"id"`
	InstrumentID   int64     `json:"instrument_id"`
	Date           time.Time `json:"date"`
	Open           int       `json:"open"`
	Low            int       `json:"low"`
	High           int       `json:"high"`
	Close          int       `json:"close"`
}
