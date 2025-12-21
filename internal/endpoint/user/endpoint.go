package user

import (
	"github.com/JeyKeyAlex/TourProject/internal/service/user"
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetUserList endpoint.Endpoint
	Create      endpoint.Endpoint
	Approve     endpoint.Endpoint
	GetUserById endpoint.Endpoint
	Delete      endpoint.Endpoint
}

func MakeEndpoints(s user.IService) Endpoints {
	return Endpoints{
		GetUserList: makeGetUserList(s),
		Create:      makeCreateUser(s),
		Approve:     makeApproveUser(s),
		GetUserById: makeGetUserById(s),
		Delete:      makeDeleteUserById(s),
	}
}
