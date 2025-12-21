package user

import (
	"context"
	"errors"
	"github.com/JeyKeyAlex/TourProject/pkg/error_templates"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/JeyKeyAlex/TourProject-proto/go-genproto/user"
)

func (s *RPCServer) Create(ctx context.Context, req *pb.CreateUserRequest) (*emptypb.Empty, error) {
	_, resp, err := s.create.ServeGRPC(ctx, req)
	if err != nil {
		var outputError *error_templates.OutputError
		if errors.As(err, &outputError) {
			return nil, outputError.GetGRPC()
		} else {
			return nil, error_templates.New(err.Error(), err, codes.Unknown, http.StatusInternalServerError).GetGRPC()
		}
	}
	return resp.(*emptypb.Empty), nil
}

func (s *RPCServer) ApproveUser(ctx context.Context, req *pb.ApproveUserRequest) (*pb.Id, error) {
	_, resp, err := s.approve.ServeGRPC(ctx, req)
	if err != nil {
		var outputError *error_templates.OutputError
		if errors.As(err, &outputError) {
			return nil, outputError.GetGRPC()
		} else {
			return nil, error_templates.New(err.Error(), err, codes.Unknown, http.StatusInternalServerError).GetGRPC()
		}
	}
	return resp.(*pb.Id), nil
}

func (s *RPCServer) DeleteUserById(ctx context.Context, req *pb.Id) (*emptypb.Empty, error) {
	_, resp, err := s.delete.ServeGRPC(ctx, req)
	if err != nil {
		var outputError *error_templates.OutputError
		if errors.As(err, &outputError) {
			return nil, outputError.GetGRPC()
		} else {
			return nil, error_templates.New(err.Error(), err, codes.Unknown, http.StatusInternalServerError).GetGRPC()
		}
	}
	return resp.(*emptypb.Empty), nil
}

func (s *RPCServer) GetUserList(ctx context.Context, req *emptypb.Empty) (*pb.GetUserListResponse, error) {
	_, resp, err := s.getUserList.ServeGRPC(ctx, req)
	if err != nil {
		var outputError *error_templates.OutputError
		if errors.As(err, &outputError) {
			return nil, outputError.GetGRPC()
		} else {
			return nil, error_templates.New(err.Error(), err, codes.Unknown, http.StatusInternalServerError).GetGRPC()
		}
	}
	return resp.(*pb.GetUserListResponse), nil
}

func (s *RPCServer) GetUserById(ctx context.Context, req *pb.Id) (*pb.GetUserByIdResponse, error) {
	_, resp, err := s.getUserById.ServeGRPC(ctx, req)
	if err != nil {
		var outputError *error_templates.OutputError
		if errors.As(err, &outputError) {
			return nil, outputError.GetGRPC()
		} else {
			return nil, error_templates.New(err.Error(), err, codes.Unknown, http.StatusInternalServerError).GetGRPC()
		}
	}
	return resp.(*pb.GetUserByIdResponse), nil
}
