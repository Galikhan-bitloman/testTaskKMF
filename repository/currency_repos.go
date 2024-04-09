package repository

import (
	"log"

	"github.com/jmoiron/sqlx"
)

type Currency struct {
	db *sqlx.DB
}

func NewCurrencyStruct(db *sqlx.DB) *Currency {
	return &Currency{
		db: db,
	}
}

func (c *Currency) CreateCurrencyRepository() {
	log.Println("here will be CreateCurrencyRepository functionality")
}
