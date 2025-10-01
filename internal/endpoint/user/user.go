package user

import (
	"context"
	"net/http"

	"github.com/JeyKeyAlex/TourProject/internal/endpoint"
	"github.com/JeyKeyAlex/TourProject/internal/entities"
	userSrv "github.com/JeyKeyAlex/TourProject/internal/service/user"
)

func GetUserListHandler(s userSrv.IService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		userList, err := s.GetUserList(ctx)
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
