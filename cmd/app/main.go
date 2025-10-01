package main

import (
	"fmt"
	"net/http"

	"github.com/JeyKeyAlex/TourProject/internal/database"
	"github.com/JeyKeyAlex/TourProject/internal/service/user"
	httpApi "github.com/JeyKeyAlex/TourProject/internal/transport/http"
)

func main() {
	// TODO init config
	// TODO init logger
	// TODO init runtime

	// database initialization

	// 1. Create DB pool
	db, err := DBinit("postgres://postgres:9512357@localhost:5432/test?sslmode=disable") // убрать в конфиги
	if err != nil {
		panic(err)
	} else {
		fmt.Println("successful connection to database")
	}
	defer db.Close()

	// 2. NewDBP creates a new database sample
	RWDBOperationer := database.NewDBP(db)

	// TODO init service
	srv := user.NewService(RWDBOperationer)
	// TODO init client (grpc, http, smtp,...)
	chiRouter := initRouter()
	httpApi.InitApi(chiRouter, srv)
	fmt.Println("starting server on port 8080")   // убрать в init
	err = http.ListenAndServe(":8080", chiRouter) // убрать в конфиги
	if err != nil {
		panic(err)
	}
	// TODO init messenger broker (Kafka, Rabbit, Nats)
	// TODO init transport
}
