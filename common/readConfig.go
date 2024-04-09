package common

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Conf struct {
	Port             string
	ConnectionString Server
}

type Server struct {
	Server   string
	Port     int
	User     string
	Password string
	Database string
}

func ReadConfig() (res Conf) {
	var AppConfig Conf
	raw, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Println("Error occured while reading config")
		return
	}
	err = json.Unmarshal(raw, &AppConfig)
	if err != nil {
		return
	}

	return AppConfig
}
