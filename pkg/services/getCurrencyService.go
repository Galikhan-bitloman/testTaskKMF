package services

import (
	"log"
	"testTaskKMF/pkg/common"
	"testTaskKMF/pkg/consts"
	"testTaskKMF/pkg/repository"
	"testTaskKMF/pkg/schema"
)

type GetCurrency struct {
	repos  repository.ICurrencyRepository
	config common.Conf
}

func NewGetCurrency(repos repository.ICurrencyRepository, config common.Conf) *GetCurrency {
	return &GetCurrency{
		repos:  repos,
		config: config,
	}
}

func (g *GetCurrency) GetCurrencyService(queryValue string) (*schema.CurrencyResponse, error) {
	var rates schema.Rates

	url, err := GetRatesUrlConstructor(queryValue, g.config.URL.GetRates)
	if err != nil {
		return nil, err
	}

	log.Println("url", *url)

	byteResp, err := MakeRequests(consts.GET, url)
	if err != nil {
		return nil, err
	}

	res, err := XmlToJsonCurrency(byteResp, &rates)
	if err != nil {
		return nil, err
	}

	return res, nil
}
