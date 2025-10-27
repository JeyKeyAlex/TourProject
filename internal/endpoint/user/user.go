package user

import (
	"context"
	userSrv "github.com/JeyKeyAlex/TourProject/internal/service/user"
	"github.com/go-kit/kit/endpoint"
)

func makeGetUserList(s userSrv.IService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		resp, err := s.GetUserList(ctx)
		if err != nil {
			return nil, err
		}

		return &resp, nil
	}
}
