package main

import (
	"log"
	"testTaskKMF/common"
	"testTaskKMF/handler"
	"testTaskKMF/repository"
	"testTaskKMF/services"
)

func main() {

	config, err := common.ReadConfig("config/config.json")

	if err != nil {
		log.Fatal(err)
	}

	db, err := repository.DBConnection(
		repository.Config{
			ConnectionString: repository.Server{
				Host:     "localhost",
				Port:     5432,
				User:     "galikhan",
				Password: "postgres",
				Database: "postgres",
				SSLMode:  "disable",
			},
		},
	)

	if err != nil {
		log.Fatal(err)
	}

	repos := repository.NewRepository(db)
	service := services.NewService(repos, config)
	handler := handler.NewHandler(service)
	routes, err := handler.InitRoutes()
	if err != nil {
		log.Fatal(err)
	}

	err = routes.Listen(":8000")
	if err != nil {
		log.Fatal(err)
	}
}
