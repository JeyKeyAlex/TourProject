package main

import (
	"context"
	"github.com/JeyKeyAlex/TourProject/internal/config"
	"github.com/JeyKeyAlex/TourProject/pkg/logger"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/JeyKeyAlex/TourProject/internal/database"
	"github.com/rs/zerolog"
)

func main() {
	// TODO init config
	appConfig, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	// TODO init profiler
	// TODO init logger
	// loggers
	var baseLogger zerolog.Logger
	var loggerCloser io.WriteCloser
	if appConfig.Log.Batch {
		baseLogger, loggerCloser, err = logger.NewDiodeLogger(os.Stdout, appConfig.Log.Level, appConfig.Log.BatchSize, appConfig.Log.BatchPollInterval)
	} else {
		baseLogger, loggerCloser, err = logger.NewLogger(os.Stdout, appConfig.Log.Level)
	}
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if loggerCloser != nil {
			err = loggerCloser.Close()
			if err != nil {
				log.Fatalf("error acquired while closing log writer: %+v", err)
			}
		}
	}()
	baseLogger = baseLogger.With().
		Str("app_version", appConfig.Version.Number).
		Str("app_build", appConfig.Version.Build).
		CallerWithSkipFrameCount(2).
		Logger()
	//apiLogger := logger.NewComponentLogger(baseLogger, "api")
	coreLogger := logger.NewComponentLogger(baseLogger, "core")
	netLogger := logger.NewComponentLogger(baseLogger, "net")

	defer func() {
		coreLogger.Info().Msg("application stopped")
	}()

	coreLogger.Info().Msg("system initialization started")

	// TODO init runtime

	// ---------  database initialization  ---------

	// 1. Create DB pool

	db, err := DBinit(appConfig.RWDB.ConnectionString) // убрать в конфиги
	if err != nil {
		panic(err)
	} else {
		coreLogger.Info().Msg("database initialization started")
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
	httpServer, httpListener := initKitHTTP(serviceEndpoints, chiRouter, listenErr, appConfig, netLogger)

	defer func() {
		err = httpListener.Close()
		if err != nil {
			// we don't really need it because we already closed it by cmux.Close()
			// netLogger.Warn().Err(err).Msgf("failed to close net.Listen %+w - %+v", err, err)
		}
	}()

	// ---------  server initialized  ---------
	runApp(httpServer, listenErr, coreLogger)
}

func runApp(httpServer *http.Server, listenErr chan error, coreLogger zerolog.Logger) {

	var shutdownCh = make(chan os.Signal, 1)
	signal.Notify(shutdownCh, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	var err error
	var runningApp = true

	for runningApp {
		select {
		// handle error channel
		case err = <-listenErr:
			if err != nil {
				coreLogger.Error().Err(err).Msg("received listener error")
				shutdownCh <- os.Kill
			}
		// handle os system signal
		case sig := <-shutdownCh:
			coreLogger.Info().Msgf("received shutdown signal: %v", sig)
			ctxTimeout, timeoutCancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
			err = httpServer.Shutdown(ctxTimeout) // may return ErrServerClosed
			defer timeoutCancelFunc()
			if err != nil {
				coreLogger.Error().Err(err).Msg("received shutdown error")
			}
			runningApp = false
		}
	}
}
