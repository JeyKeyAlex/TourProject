package user

import (
	"context"
	"errors"
	"github.com/JeyKeyAlex/TourProject/internal/entities"
	userSrv "github.com/JeyKeyAlex/TourProject/internal/service/user"
	pkgErr "github.com/JeyKeyAlex/TourProject/pkg/errors"
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

func makeCreateUser(s userSrv.IService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		serviceLogger := s.GetLogger().With().Str("func", "makeCreateUser").Logger()
		serviceLogger.Info().Msg("calling s.createUser")

		req, err := helpers.CastRequest[*entities.CreateUserRequest](request)
		if err != nil {
			serviceLogger.Error().Stack().Err(error_templates.ErrorDetailFromError(err)).Msg(pkgErr.FailedCastRequest)
			return nil, err
		}

		err = s.CreateUser(ctx, req)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}
}

func makeApproveUser(s userSrv.IService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		serviceLogger := s.GetLogger().With().Str("func", "makeApproveUser").Logger()
		serviceLogger.Info().Msg("calling s.makeApproveUser")

		email, ok := request.(string)
		if !ok {
			err := errors.New("email must be a string")
			serviceLogger.Error().Stack().Err(error_templates.ErrorDetailFromError(err)).Msg(pkgErr.FailedCastRequest)
			return nil, err
		}

		id, err := s.ApproveUser(ctx, email)
		if err != nil {
			return nil, err
		}

		return &id, nil
	}
}

func makeGetUserById(s userSrv.IService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		serviceLogger := s.GetLogger().With().Str("func", "makeCreateUser").Logger()
		serviceLogger.Info().Msg("calling s.createUser")

		userId, ok := request.(string)
		if !ok {
			err := errors.New("userid must be a string")
			serviceLogger.Error().Stack().Err(error_templates.ErrorDetailFromError(err)).Msg(pkgErr.FailedCastRequest)
			return nil, err
		}

		user, err := s.GetUserById(ctx, userId)
		if err != nil {
			return nil, err
		}

		return &user, nil
	}
}
