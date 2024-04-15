package repository

import "github.com/jmoiron/sqlx"

type ICurrencyRepository interface {
	CreateCurrencyRepository()
}

type Repository struct {
	ICurrencyRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		ICurrencyRepository: NewCurrencyStruct(db),
	}
}
