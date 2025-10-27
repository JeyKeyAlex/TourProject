package main

import (
	"context"
	"github.com/JeyKeyAlex/TourProject/internal/config"
	"github.com/rs/zerolog"
	"net"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/JeyKeyAlex/TourProject/internal/database"
	"github.com/JeyKeyAlex/TourProject/internal/endpoint"
	userEp "github.com/JeyKeyAlex/TourProject/internal/endpoint/user"
	srvUser "github.com/JeyKeyAlex/TourProject/internal/service/user"
	tpHTTP "github.com/JeyKeyAlex/TourProject/internal/transport/http"
	custumMiddlware "github.com/JeyKeyAlex/TourProject/internal/transport/http/middleware"
	tpHTTPUser "github.com/JeyKeyAlex/TourProject/internal/transport/http/user"
)

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
	rwdbOperationer database.RWDBOperationer,
) endpoint.ServiceEndpoints {
	userSrv := srvUser.NewService(rwdbOperationer)
	return endpoint.ServiceEndpoints{
		UserEP: userEp.MakeEndpoints(userSrv),
	}
}

func initHTTPRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.NoCache)
	r.Use(middleware.RealIP)
	r.Use(middleware.RequestID)
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
