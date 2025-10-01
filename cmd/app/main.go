package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/JeyKeyAlex/TourProject/internal/database"
)

func main() {
	// TODO init config
	// TODO init logger
	// TODO init runtime

	// ---------  database initialization  ---------

	// 1. Create DB pool

	db, err := DBinit("postgres://postgres:9512357@localhost:5432/test?sslmode=disable") // убрать в конфиги
	if err != nil {
		panic(err)
	} else {
		fmt.Println("successful connection to database")
	}
	defer db.Close()

	// 2. NewRWDBOperationer creates a new database sample

	RWDBOperationer := database.NewRWDBOperationer(db)

	// ---------  database initialized ---------

	// ---------  service initialization  ---------

	serviceEndpoints := initEndpoints(RWDBOperationer)

	// ---------  service initialized  ---------

	// TODO init client (grpc, http, smtp,...)
	// TODO init messenger broker (Kafka, Rabbit, Nats)

	// ---------  server initialization  ---------
	//httpApi.InitApi(chiRouter, srv)
	//fmt.Println("starting server on port 8080")   // убрать в init
	//err = http.ListenAndServe(":8080", chiRouter) // убрать в конфиги
	//if err != nil {
	//	panic(err)
	//}
	chiRouter := initHTTPRouter()
	// TODO healthChecker

	listenErr := make(chan error, 1)
	httpServer, httpListener := initKitHTTP(serviceEndpoints, chiRouter, listenErr)

	defer func() {
		err = httpListener.Close()
		if err != nil {
			// we don't really need it because we already closed it by cmux.Close()
			// netLogger.Warn().Err(err).Msgf("failed to close net.Listen %+w - %+v", err, err)
		}
	}()

	// ---------  server initialized  ---------
	runApp(httpServer, listenErr)
}

func runApp(httpServer *http.Server, listenErr chan error) {

	var shutdownCh = make(chan os.Signal, 1)
	signal.Notify(shutdownCh, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	var err error
	var runningApp = true

	for runningApp {
		select {
		// handle error channel
		case err = <-listenErr:
			if err != nil {
				fmt.Println("received listener error")
				shutdownCh <- os.Kill
			}
		// handle os system signal
		case sig := <-shutdownCh:
			fmt.Printf("shutdown signal received: %s", sig.String())
			ctxTimeout, timeoutCancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
			err = httpServer.Shutdown(ctxTimeout) // may return ErrServerClosed
			defer timeoutCancelFunc()
			fmt.Println("received http Shutdown error")
			runningApp = false
		}
	}
}
