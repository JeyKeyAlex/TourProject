package user

import (
	"github.com/JeyKeyAlex/TourProject/internal/entities"
	"net/http"

	"github.com/JeyKeyAlex/TourProject/internal/database"
	"github.com/JeyKeyAlex/TourProject/internal/endpoint"
)

func GetUserListHandler(s database.RWDBOperationer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userList, err := s.GetUserList()
		if err != nil {
			endpoint.WriteJSON(w, http.StatusBadRequest, err)
			return
		}
		if userList == nil {
			endpoint.WriteJSON(w, http.StatusOK, []entities.User{})
		} else {
			endpoint.WriteJSON(w, http.StatusOK, userList)
		}
	}
}
