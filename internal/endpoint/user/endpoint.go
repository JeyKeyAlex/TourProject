package user

import (
	"github.com/JeyKeyAlex/TourProject/internal/service/user"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetUser     endpoint.Endpoint
	GetUserList endpoint.Endpoint
	Create      endpoint.Endpoint
	Approve     endpoint.Endpoint
	Update      endpoint.Endpoint
	Delete      endpoint.Endpoint
}

func MakeEndpoints(s user.IService) Endpoints {
	return Endpoints{
		GetUserList: makeGetUserList(s),
		GetUser:     makeGetUser(s),
		Create:      makeCreate(s),
		Approve:     makeApprove(s),
		Update:      makeUpdate(s),
		Delete:      makeDelete(s),
	}
}
