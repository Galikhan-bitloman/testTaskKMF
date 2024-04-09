package services

import (
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"testTaskKMF/repository"
	"testTaskKMF/schema"
)

type GetCurrency struct {
	repos repository.ICurrencyRepository
}

func NewGetCurrency(repos repository.ICurrencyRepository) *GetCurrency {
	return &GetCurrency{
		repos: repos,
	}
}

func (g *GetCurrency) GetCurrencyService(queryValue string) (*schema.Rates, error) {
	var res schema.Rates

	log.Println("bus logic for get currency")

	url, err := GetRatesUrlConstructor(queryValue)

	log.Println("url", *url)
	if err != nil {
		log.Println(err)
	}

	request, err := http.NewRequest("GET", *url, nil)

	if err != nil {
		return nil, err
	}

	client := http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	responseBody, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	err = xml.Unmarshal(responseBody, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil

}
