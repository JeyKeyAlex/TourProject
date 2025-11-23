package user

import (
	"github.com/JeyKeyAlex/TourProject/internal/endpoint/user"
	"github.com/JeyKeyAlex/TourProject/internal/transport/http/common"
	custumMiddlware "github.com/JeyKeyAlex/TourProject/internal/transport/http/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	kithttp "github.com/go-kit/kit/transport/http"
	"net/http"
)

func NewServer(endpoints user.Endpoints, options []kithttp.ServerOption) http.Handler {
	r := chi.NewRouter()

	options = append(options, kithttp.ServerErrorEncoder(common.EncodeErrorResponse))

	r.Use(middleware.NoCache)
	r.Use(middleware.RealIP)
	r.Use(custumMiddlware.RequestID)
	r.Use(middleware.Recoverer)
	r.Use(middleware.StripSlashes)

	r.Get("/users", kithttp.NewServer(endpoints.GetUserList, decodeEmptyRequest, kithttp.EncodeJSONResponse, options...).ServeHTTP)
	r.Post("/users", kithttp.NewServer(endpoints.CreateUser, decodeCreateUserRequest, kithttp.EncodeJSONResponse, options...).ServeHTTP)

	return r
}
