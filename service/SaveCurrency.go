package service

// import (
// 	"encoding/xml"
// 	"io"
// 	"log"
// 	"net/http"
// 	"sync"
// 	"testTaskKMF/common"
// 	"testTaskKMF/schema"
// )

// func SaveCurrency(w http.ResponseWriter, r *http.Request) {
// 	apiUrl := "https://nationalbank.kz/rss/get_rates.cfm?fdate=15.04.2022"

// 	request, err := http.NewRequest("GET", apiUrl, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	client := http.Client{}
// 	response, err := client.Do(request)

// 	responseBody, err := io.ReadAll(response.Body)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	var rates schema.Rates
// 	err = xml.Unmarshal(responseBody, &rates)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	var wg sync.WaitGroup
// 	for _, i := range rates.Items {
// 		wg.Add(1)
// 		i := i
// 		go func(d schema.RateItem) {
// 			common.SaveData(&i, rates.Date)
// 			wg.Done()
// 		}(i)
// 	}
// 	wg.Wait()
// }
