package user

import (
	"net/http"

	"github.com/JeyKeyAlex/TourProject/internal/endpoint/user"
	"github.com/JeyKeyAlex/TourProject/internal/transport/http/common"
	"github.com/go-chi/chi/v5"
	kithttp "github.com/go-kit/kit/transport/http"
)

func NewServer(endpoints user.Endpoints, options []kithttp.ServerOption) http.Handler {
	r := chi.NewRouter()

	options = append(options, kithttp.ServerErrorEncoder(common.EncodeErrorResponse))

	r.Get("/users", kithttp.NewServer(endpoints.GetUserList, decodeEmptyRequest, kithttp.EncodeJSONResponse, options...).ServeHTTP)
	r.Post("/users", kithttp.NewServer(endpoints.CreateUser, decodeCreateUserRequest, kithttp.EncodeJSONResponse, options...).ServeHTTP)
	r.Post("/users/approve/{email}", kithttp.NewServer(endpoints.ApproveUser, decodeApproveUserRequest, kithttp.EncodeJSONResponse, options...).ServeHTTP)
	r.Get("/users/{id}", kithttp.NewServer(endpoints.GetUserById, decodeGetUserByIdRequest, kithttp.EncodeJSONResponse, options...).ServeHTTP)
	r.Delete("/users/{id}", kithttp.NewServer(endpoints.DeleteUserById, decodeDeleteUserByIdRequest, kithttp.EncodeJSONResponse, options...).ServeHTTP)

	return r
}
