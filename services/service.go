package services

import (
	"testTaskKMF/repository"
	"testTaskKMF/schema"
)

type ICurrencyService interface {
	GetCurrencyService(queryValue string) (*schema.Rates, error)
}

type Service struct {
	ICurrencyService
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		ICurrencyService: NewGetCurrency(repos),
	}
}
