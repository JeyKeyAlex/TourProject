package user

import (
	"context"
	"errors"
	"github.com/JeyKeyAlex/TourProject/internal/convert"

	userSrv "github.com/JeyKeyAlex/TourProject/internal/service/user"
	"github.com/JeyKeyAlex/TourProject/internal/transport/http/middleware"
	"github.com/JeyKeyAlex/TourProject/pkg/error_templates"
	pkgErr "github.com/JeyKeyAlex/TourProject/pkg/errors"
	"github.com/JeyKeyAlex/TourProject/pkg/helpers/validate"

	"github.com/go-kit/kit/endpoint"

	pb "github.com/JeyKeyAlex/TourProject-proto/go-genproto/user"
)

func makeGetUserList(s userSrv.IService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		reqID, ctx := middleware.GetRequestID(ctx)
		serviceLogger := s.GetLogger().With().Str("func", "makeGetUserList").Str("request_id", reqID).Logger()
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
		reqID, ctx := middleware.GetRequestID(ctx)
		serviceLogger := s.GetLogger().With().Str("func", "makeCreateUser").Str("request_id", reqID).Logger()
		serviceLogger.Info().Msg("calling s.createUser")

		//req, err := validate.CastValidateRequest[*pb.CreateRequest](s.GetValidator(), request)
		//if err != nil {
		//	serviceLogger.Error().Stack().Err(error_templates.ErrorDetailFromError(err)).Msg(pkgErr.FailedCastRequest)
		//	return nil, err
		//}
		//
		//err = s.CreateUser(ctx, req)
		//if err != nil {
		//	return nil, err
		//}

		return nil, nil
	}
}

func makeApproveUser(s userSrv.IService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		reqID, ctx := middleware.GetRequestID(ctx)
		serviceLogger := s.GetLogger().With().Str("func", "makeApproveUser").Str("request_id", reqID).Logger()
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
		reqID, ctx := middleware.GetRequestID(ctx)
		serviceLogger := s.GetLogger().With().Str("func", "makeGetUserById").Str("request_id", reqID).Logger()
		serviceLogger.Info().Msg("calling s.GetUserById")

		req, err := validate.CastValidateRequest[*pb.IdMessage](s.GetValidator(), request)
		if err != nil {
			serviceLogger.Error().Stack().Err(error_templates.ErrorDetailFromError(err)).Msg(pkgErr.FailedCastRequest)
			return nil, err
		}

		user, err := s.GetUserById(ctx, req.Id)
		if err != nil {
			return nil, err
		}

		protoUser, err := convert.GetUserEntityToEntry(user)
		if err != nil {
			return nil, err
		}

		return &pb.GetResponse{
			User: protoUser,
		}, nil
	}
}

func makeDeleteUserById(s userSrv.IService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		reqID, ctx := middleware.GetRequestID(ctx)
		serviceLogger := s.GetLogger().With().Str("func", "makeDeleteUserById").Str("request_id", reqID).Logger()
		serviceLogger.Info().Msg("calling s.Delete")

		req, err := validate.CastValidateRequest[*pb.IdMessage](s.GetValidator(), request)
		if err != nil {
			serviceLogger.Error().Stack().Err(error_templates.ErrorDetailFromError(err)).Msg(pkgErr.FailedCastRequest)
			return nil, err
		}

		err = s.DeleteUserById(ctx, req.Id)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}
}
