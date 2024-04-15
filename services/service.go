package services

import (
	"testTaskKMF/common"
	"testTaskKMF/repository"
	"testTaskKMF/schema"
)

type ICurrencyService interface {
	GetCurrencyService(queryValue string) (*schema.CurrencyResponse, error)
}

type Service struct {
	ICurrencyService
}

func NewService(repos *repository.Repository, config *common.Conf) *Service {
	return &Service{
		ICurrencyService: NewGetCurrency(repos, *config),
	}
}
