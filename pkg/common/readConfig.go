package common

import (
	"encoding/json"
	"log"
	"os"
)

type Conf struct {
	Port             string `json:"port"`
	ConnectionString Server
	URL              Urls
}

type Urls struct {
	GetRates string `json:"getRates"`
}

type Server struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

func ReadConfig(path string) (*Conf, error) {
	var AppConfig Conf
	raw, err := os.ReadFile(path)

	if err != nil {
		log.Println("Error occured while reading config")
		return nil, err
	}
	err = json.Unmarshal(raw, &AppConfig)
	if err != nil {
		return nil, err
	}

	return &AppConfig, err
}
