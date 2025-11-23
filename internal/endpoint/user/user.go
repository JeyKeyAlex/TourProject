package user

import (
	"context"
	"github.com/JeyKeyAlex/TourProject/internal/entities"
	userSrv "github.com/JeyKeyAlex/TourProject/internal/service/user"
	"github.com/JeyKeyAlex/TourProject/pkg/errors"
	"github.com/JeyKeyAlex/TourProject/pkg/helpers"
	"github.com/go-kit/kit/endpoint"

	"github.com/JeyKeyAlex/TourProject/pkg/error_templates"
)

func makeGetUserList(s userSrv.IService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		serviceLogger := s.GetLogger().With().Str("func", "makeGetUserList").Logger()
		serviceLogger.Info().Msg("calling s.getUserList")

		resp, err := s.GetUserList(ctx)
		if err != nil {
			return nil, err
		}

		return &resp, nil
	}
}

func CreateUser(s userSrv.IService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		serviceLogger := s.GetLogger().With().Str("func", "CreateUser").Logger()
		serviceLogger.Info().Msg("calling s.createUser")

		req, err := helpers.CastRequest[*entities.CreateUserRequest](request)
		if err != nil {
			serviceLogger.Error().Stack().Err(error_templates.ErrorDetailFromError(err)).Msg(errors.FailedCastRequest)
			return nil, err
		}

		id, err := s.CreateUser(ctx, req)
		if err != nil {
			return nil, err
		}

		return &id, nil
	}
}
