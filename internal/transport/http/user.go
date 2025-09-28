package http

import (
	"github.com/go-chi/chi/v5"

	"github.com/JeyKeyAlex/TourProject/internal/endpoint/user"
)

func InitApi(router *chi.Mux) {
	router.Get("/user", user.GetUserListHandler)
}
