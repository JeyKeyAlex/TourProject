package main

import (
	"context"
	"github.com/rs/zerolog"
	"net"
	"net/http"
	"runtime"
	"runtime/debug"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/jackc/pgx/v5/pgxpool"
	goRedis "github.com/redis/go-redis/v9"
	googlegrpc "google.golang.org/grpc"

	"github.com/JeyKeyAlex/TourProject/internal/config"
	"github.com/JeyKeyAlex/TourProject/internal/database/postgreSql"
	"github.com/JeyKeyAlex/TourProject/internal/database/redis"
	"github.com/JeyKeyAlex/TourProject/internal/endpoint"
	userEp "github.com/JeyKeyAlex/TourProject/internal/endpoint/user"
	srvUser "github.com/JeyKeyAlex/TourProject/internal/service/user"
	tpGRPC "github.com/JeyKeyAlex/TourProject/internal/transport/grpc"
	tpGRPCUser "github.com/JeyKeyAlex/TourProject/internal/transport/grpc/user"
	tpHTTP "github.com/JeyKeyAlex/TourProject/internal/transport/http"
	custumMiddlware "github.com/JeyKeyAlex/TourProject/internal/transport/http/middleware"
	tpHTTPUser "github.com/JeyKeyAlex/TourProject/internal/transport/http/user"

	pbUser "github.com/JeyKeyAlex/TourProject-proto/go-genproto/user"
)

func initRuntime(useCPUs, maxThreads int, logger *zerolog.Logger) {
	if useCPUs == 0 {
		useCPUs = runtime.NumCPU()
		runtime.GOMAXPROCS(useCPUs)
	} else {
		runtime.GOMAXPROCS(useCPUs)
	}
	logger.Info().Msgf("num of CPUs: %d", useCPUs)

	if maxThreads == 0 {
		maxThreads = 10000
	}
	debug.SetMaxThreads(maxThreads)
	logger.Info().Msgf("max threads: %d", maxThreads)
}

func DBinit(connectionString string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		return nil, err
	}

	err = pool.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	return pool, nil
}

func initEndpoints(
	rwdbOperationer postgreSql.RWDBOperationer,
	redisDB redis.Redis,
	logger *zerolog.Logger,
	appConfig *config.Configuration,
) endpoint.ServiceEndpoints {
	userSrv := srvUser.NewService(rwdbOperationer, redisDB, logger, appConfig)
	return endpoint.ServiceEndpoints{
		UserEP: userEp.MakeEndpoints(userSrv),
	}
}

func initHTTPRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.NoCache)
	r.Use(middleware.RealIP)
	r.Use(custumMiddlware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middleware.StripSlashes)

	pongResponse := []byte("pong")
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write(pongResponse)
	})

	return r
}

func initKitHTTP(endpoints endpoint.ServiceEndpoints, router *chi.Mux, listenErr chan error, cfg *config.Configuration, netLogger zerolog.Logger) (*http.Server, net.Listener) {
	var serverOptions []kithttp.ServerOption
	router.Mount("/UserService/", tpHTTPUser.NewServer(endpoints.UserEP, serverOptions))

	httpServer := &http.Server{
		Handler:      router,
		TLSConfig:    nil,
		ReadTimeout:  cfg.HTTP.ReadTimeout,
		WriteTimeout: cfg.HTTP.WriteTimeout,
		IdleTimeout:  cfg.HTTP.IdleTimeout,
	}

	l, err := net.Listen(cfg.HTTP.Network, cfg.HTTP.Address)
	if err != nil {
		netLogger.Fatal().Err(err).Msg("failed to init net.Listen for http")
	} else {
		netLogger.Info().Msg("successful net.Listen for http init")
	}

	go tpHTTP.RunHTTPServer(httpServer, l, listenErr, netLogger)
	time.Sleep(10 * time.Millisecond)

	return httpServer, l
}

func initKitGRPC(appConfig *config.Configuration, endpoints endpoint.ServiceEndpoints, netLogger zerolog.Logger, listenErr chan error) (*googlegrpc.Server, net.Listener) {
	var serverOptions []kitgrpc.ServerOption

	grpcUserServer := tpGRPCUser.NewServer(endpoints.UserEP, serverOptions)

	grpcServer := googlegrpc.NewServer(
		googlegrpc.MaxRecvMsgSize(appConfig.GRPC.MaxRequestBodySize),
		googlegrpc.MaxSendMsgSize(appConfig.GRPC.MaxRequestBodySize),
	)

	pbUser.RegisterUserServiceServer(grpcServer, grpcUserServer)

	l, err := net.Listen(appConfig.GRPC.Network, appConfig.GRPC.Address)
	if err != nil {
		netLogger.Fatal().Err(err).Msg("failed to init net.Listen for grpc")
	} else {
		netLogger.Info().Msg("successful net.Listen for grpc init")
	}

	go tpGRPC.RunGRPCServer(grpcServer, l, netLogger, listenErr)
	time.Sleep(10 * time.Millisecond)
	return grpcServer, l
}

func initRedisConnection(appConfig *config.Configuration) (*goRedis.Client, error) {
	opts, err := goRedis.ParseURL(appConfig.Redis.ConnectionString)
	if err != nil {
		return nil, err
	}
	rds := goRedis.NewClient(opts)

	_, err = rds.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return rds, nil
}
