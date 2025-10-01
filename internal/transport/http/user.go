package http

import (
	"github.com/go-chi/chi/v5"

	"github.com/JeyKeyAlex/TourProject/internal/endpoint/user"
	userSrv "github.com/JeyKeyAlex/TourProject/internal/service/user"
)

func InitApi(router *chi.Mux, srv userSrv.IService) {
	router.Get("/users", user.GetUserListHandler(srv))
}
