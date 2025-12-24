package user

import (
	"github.com/JeyKeyAlex/TourProject/internal/service/user"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	Create      endpoint.Endpoint
	Approve     endpoint.Endpoint
	Delete      endpoint.Endpoint
	Update      endpoint.Endpoint
	GetUser     endpoint.Endpoint
	GetUserList endpoint.Endpoint
}

func MakeEndpoints(s user.IService) Endpoints {
	return Endpoints{
		Create:      makeCreate(s),
		Approve:     makeApprove(s),
		Update:      makeApprove(s),
		Delete:      makeDelete(s),
		GetUserList: makeGetUserList(s),
		GetUser:     makeGetUser(s),
	}
}
