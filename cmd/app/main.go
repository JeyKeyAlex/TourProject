package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/JeyKeyAlex/TourProject/internal/config"
	"github.com/JeyKeyAlex/TourProject/internal/database/postgreSql"
	"github.com/JeyKeyAlex/TourProject/internal/database/redis"
	"github.com/JeyKeyAlex/TourProject/pkg/logger"

	goRedis "github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"

	"buf.build/go/protovalidate"
)

func main() {
	appConfig, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	// TODO init profiler

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

	initRuntime(appConfig.RunTime.UseCPUs, appConfig.RunTime.MaxThreads, &coreLogger)

	rwdb, err := DBinit(appConfig.RWDB.ConnectionString)
	if err != nil {
		coreLogger.Fatal().Err(err).Msg("failed to establish a connection with the Read/Write database")
	} else {
		coreLogger.Info().Msg("database initialization started")
	}
	defer rwdb.Close()

	rwDb := postgreSql.New(rwdb, &appConfig.RWDB)

	rds, err := initRedisConnection(appConfig)
	if err != nil {
		coreLogger.Fatal().Err(err).Msg("failed to establish a connection with the redis")
	} else {
		coreLogger.Info().Msg("successful connection with the redis")
	}
	defer func(rds *goRedis.Client) {
		err = rds.Close()
		if err != nil {
			coreLogger.Error().Msg("failed to close the redis connection")
		}
	}(rds)

	redisDB, err := redis.New(rds)

	validator, err := protovalidate.New()
	serviceEndpoints := initEndpoints(rwDb, redisDB, validator, &netLogger, appConfig)

	// TODO init client (grpc, http, smtp,...)
	// TODO init messenger broker (Kafka, Rabbit, Nats)

	chiRouter := initHTTPRouter()
	// TODO healthChecker

	listenErr := make(chan error, 1)
	httpServer, httpListener := initKitHTTP(serviceEndpoints, chiRouter, listenErr, appConfig, netLogger)

	defer func() {
		err = httpListener.Close()
		if err != nil {
			coreLogger.Fatal().Err(err).Msg("error closing http listener")
		}
	}()

	grpcServer, grpcListener := initKitGRPC(appConfig, serviceEndpoints, netLogger, listenErr)
	defer func() {
		err = grpcListener.Close()
		if err != nil {
			coreLogger.Fatal().Err(err).Msg("error closing grpc listener")
		}
	}()

	runApp(grpcServer, httpServer, listenErr, coreLogger)
}

func runApp(grpcServer *grpc.Server, httpServer *http.Server, listenErr chan error, coreLogger zerolog.Logger) {

	var shutdownCh = make(chan os.Signal, 1)
	signal.Notify(shutdownCh, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	var err error

	for {
		select {
		case err = <-listenErr:
			if err != nil {
				coreLogger.Error().Err(err).Msg("received listener error")
				shutdownCh <- os.Kill
			}
		case sig := <-shutdownCh:
			coreLogger.Info().Msgf("received shutdown signal: %v", sig)
			ctxTimeout, timeoutCancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
			err = httpServer.Shutdown(ctxTimeout) // may return ErrServerClosed
			timeoutCancelFunc()
			if err != nil && err != http.ErrServerClosed {
				coreLogger.Error().Err(err).Msg("received shutdown error")
			}
			grpcServer.GracefulStop()
			coreLogger.Info().Msg("server loop stopped")
			return
		}
	}
}
