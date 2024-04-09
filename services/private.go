package services

import (
	"log"
	"net/url"
	"time"
)

func GetRatesUrlConstructor(queryValue string) (*string, error) {
	existingURL := "https://nationalbank.kz/rss/get_rates.cfm?fdate=param"

	parsedURL, err := url.Parse(existingURL)

	if err != nil {
		log.Println("Error parsing URL:", err)
		return nil, err
	}
	val, _ := time.Parse("02.02.2006", queryValue)

	queryParams := parsedURL.Query()
	queryParams.Add("fdate", "param")
	queryParams.Set("fdate", val.Format("02.02.2006"))

	parsedURL.RawQuery = queryParams.Encode()

	url := parsedURL.String()
	return &url, nil
}

// currentTime := time.Now().Format("15.04.2024")

// 	existingURL := "https://nationalbank.kz/rss/get_rates.cfm?fdate=param"

// 	parsedURL, err := url.Parse(existingURL)
// 	if err != nil {
// 		log.Println("Error parsing URL:", err)
// 		return nil, err
// 	}

// 	queryParams := parsedURL.Query()
// 	queryParams.Add("fdate", "param")
// 	// queryParams.Set("fdate", formattedTime)
// 	log.Println("higher", currentTime)
// 	// log.Println("queryParams", formattedTime)
// 	parsedURL.RawQuery = queryParams.Encode()
// 	return parsedURL, nil
