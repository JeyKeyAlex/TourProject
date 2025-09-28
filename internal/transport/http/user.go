package http

import (
	"github.com/go-chi/chi/v5"

	"github.com/JeyKeyAlex/TourProject/internal/database"
	"github.com/JeyKeyAlex/TourProject/internal/endpoint/user"
)

func InitApi(router *chi.Mux, db database.RWDBOperationer) {
	router.Get("/user", user.GetUserListHandler(db))
}
