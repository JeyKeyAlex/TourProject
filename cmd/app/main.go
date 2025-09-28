package main

import (
	"fmt"
	httpApi "github.com/JeyKeyAlex/TourProject/internal/transport/http"
	"net/http"
)

func main() {
	// TODO init config
	// TODO init logger
	// TODO init runtime
	// TODO init database
	db, err := DBinit("postgres://postgres:9512357@localhost:5432/test?sslmode=disable")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("successful connection to database")
	}
	defer db.Close()

	// TODO init service
	// TODO init client (grpc, http, smtp,...)
	chiRouter := initRouter()
	httpApi.InitApi(chiRouter)
	fmt.Println("starting server on port 8080")
	err = http.ListenAndServe(":8080", chiRouter)
	if err != nil {
		panic(err)
	}
	// TODO init messenger broker (Kafka, Rabbit, Nats)
	// TODO init transport
}
