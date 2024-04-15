package services

import (
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"testTaskKMF/consts"
	"testTaskKMF/schema"
	"time"
)

func GetRatesUrlConstructor(queryValue string, urlPath string) (*string, error) {

	parsedURL, err := url.Parse(urlPath)

	if err != nil {
		log.Println("Error parsing URL:", err)
		return nil, err
	}

	switch queryValue {
	case "":
		today := todayDate()

		url := urlConstructor(parsedURL, today)

		return &url, nil
	default:

		url := urlConstructor(parsedURL, queryValue)
		return &url, nil
	}
}

func urlConstructor(parsedURL *url.URL, valTime string) string {
	log.Println("valTime", valTime)
	queryParams := parsedURL.Query()
	queryParams.Add("fdate", "param")
	queryParams.Set("fdate", valTime)
	parsedURL.RawQuery = queryParams.Encode()

	url := parsedURL.String()

	return url
}

func XmlToJsonCurrency(resp []byte, resModel *schema.Rates) (*schema.CurrencyResponse, error) {

	var res schema.CurrencyResponse
	var resItem schema.CurrencyData

	err := xml.Unmarshal(resp, &resModel)

	if err != nil {
		return nil, err
	}

	for _, item := range resModel.Items {
		resItem.Change = item.Change
		resItem.ExchangeRate = item.Description
		resItem.FullName = item.Fullname
		resItem.Index = item.Index

		res.Data = append(res.Data, resItem)
		res.Date = resModel.Date
	}

	return &res, nil
}

func MakeRequests(method string, url *string) ([]byte, error) {

	upper := strings.ToUpper(method)
	switch upper {
	case consts.GET:
		request, err := http.NewRequest(consts.GET, *url, nil)

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
		return responseBody, nil
	}
	return nil, nil
}

func todayDate() string {
	return time.Now().Format("02.01.2006")
}
