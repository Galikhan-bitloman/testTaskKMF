package service

// import (
// 	"log"
// 	"net/http"
// 	"net/url"
// 	"testTaskKMF/common"
// )

// func GetCurrency(w http.ResponseWriter, r *http.Request) {
// 	apiUrl := "https://nationalbank.kz/rss/get_rates.cfm?fdate=15.04.2022"

// 	parsedURL, _ := url.Parse(apiUrl)

// 	queryParams := parsedURL.Query()

// 	_, err := common.SelectDataByID(queryParams.Get("fdate"))

// 	if err.Error != nil {
// 		log.Fatal(err)
// 	}

// }
