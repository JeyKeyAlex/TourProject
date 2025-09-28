package main

import (
	"fmt"
	"net/http"

	"github.com/JeyKeyAlex/TourProject/internal/database"
	httpApi "github.com/JeyKeyAlex/TourProject/internal/transport/http"
)

func main() {
	// TODO init config
	// TODO init logger
	// TODO init runtime
	// TODO init database

	// Create DB pool
	db, err := DBinit("postgres://postgres:9512357@localhost:5432/test?sslmode=disable") // убрать в конфиги
	if err != nil {
		panic(err)
	} else {
		fmt.Println("successful connection to database")
	}
	defer db.Close()

	// NewDBP creates a new database instance
	dbInstance := database.NewDBP(db)

	// TODO init service
	// TODO init client (grpc, http, smtp,...)
	chiRouter := initRouter()
	httpApi.InitApi(chiRouter, dbInstance)
	fmt.Println("starting server on port 8080")   // убрать в init
	err = http.ListenAndServe(":8080", chiRouter) // убрать в конфиги
	if err != nil {
		panic(err)
	}
	// TODO init messenger broker (Kafka, Rabbit, Nats)
	// TODO init transport
}
