package main

import (
	"log"
	"testTaskKMF/handler"
	"testTaskKMF/repository"
	"testTaskKMF/services"
)

func main() {

	// // c := common.ReadConfig()

	// // r := mux.NewRouter()
	// // r.HandleFunc("/currency", service.SaveCurrency)
	// // r.HandleFunc("/currency/{date}", service.GetCurrency)
	// // r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
	// // 	httpSwagger.URL("http://localhost:8080/swagger/swagger.json"), // The URL to fetch the swagger.json file
	// // ))
	// // http.Handle("/", r)
	// // fmt.Println("Server listening on :8080...")
	// // err := http.ListenAndServe(c.Port, nil)
	// // if err != nil {
	// // 	panic(err)
	// }

	db, err := repository.DBConnection(
		repository.Config{
			ConnectionString: repository.Server{
				Server:   "ms",
				Port:     1433,
				User:     "kursUser",
				Password: "kursPswd",
				Database: "TEST",
			},
		},
	)

	repos := repository.NewRepository(db)
	service := services.NewService(repos)
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
