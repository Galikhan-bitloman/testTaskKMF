package schema

import (
	"encoding/xml"
)

type Rates struct {
	XMLName     xml.Name   `xml:"rates"`
	Generator   string     `xml:"generator"`
	Title       string     `xml:"title"`
	Link        string     `xml:"link"`
	Description string     `xml:"description"`
	Copyright   string     `xml:"copyright"`
	Date        string     `xml:"date"`
	Items       []RateItem `xml:"item"`
}

type RateItem struct {
	XMLName     xml.Name `xml:"item"`
	Fullname    string   `xml:"fullname"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Quant       string   `xml:"quant"`
	Index       string   `xml:"index"`
	Change      string   `xml:"change"`
}

type CurrencyRequest struct {
	Date string `json:"date, omitempty"`
}

type CurrencyData struct {
	FullName     string `json:"fullname"`
	ExchangeRate string `json:"exchangeRate"`
	Index        string `json:"index"`
	Change       string `json:"change"`
}

type CurrencyResponse struct {
	Data []CurrencyData `json:"data"`
	Date string         `json:"date"`
}
