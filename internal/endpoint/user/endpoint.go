package user

import (
	"github.com/JeyKeyAlex/TourProject/internal/service/user"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetUserList endpoint.Endpoint
	CreateUser  endpoint.Endpoint
	GetUserById endpoint.Endpoint
}

func MakeEndpoints(s user.IService) Endpoints {
	return Endpoints{
		GetUserList: makeGetUserList(s),
		CreateUser:  makeCreateUser(s),
		GetUserById: makeGetUserById(s),
	}
}
