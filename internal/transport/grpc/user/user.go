package user

import (
	"context"
	"errors"
	"net/http"

	"github.com/JeyKeyAlex/TourProject/pkg/error_templates"

	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/JeyKeyAlex/TestProject-genproto/user"
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

func (s *RPCServer) Approve(ctx context.Context, req *pb.ApproveUserRequest) (*pb.IdMessage, error) {
	_, resp, err := s.approve.ServeGRPC(ctx, req)
	if err != nil {
		var outputError *error_templates.OutputError
		if errors.As(err, &outputError) {
			return nil, outputError.GetGRPC()
		} else {
			return nil, error_templates.New(err.Error(), err, codes.Unknown, http.StatusInternalServerError).GetGRPC()
		}
	}
	return resp.(*pb.IdMessage), nil
}

func (s *RPCServer) Update(ctx context.Context, req *pb.UpdateUserRequest) (*pb.IdMessage, error) {
	_, resp, err := s.update.ServeGRPC(ctx, req)
	if err != nil {
		var outputError *error_templates.OutputError
		if errors.As(err, &outputError) {
			return nil, outputError.GetGRPC()
		} else {
			return nil, error_templates.New(err.Error(), err, codes.Unknown, http.StatusInternalServerError).GetGRPC()
		}
	}
	return resp.(*pb.IdMessage), nil
}

func (s *RPCServer) Delete(ctx context.Context, req *pb.IdMessage) (*emptypb.Empty, error) {
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

func (s *RPCServer) GetUser(ctx context.Context, req *pb.IdMessage) (*pb.GetUserResponse, error) {
	_, resp, err := s.getUser.ServeGRPC(ctx, req)
	if err != nil {
		var outputError *error_templates.OutputError
		if errors.As(err, &outputError) {
			return nil, outputError.GetGRPC()
		} else {
			return nil, error_templates.New(err.Error(), err, codes.Unknown, http.StatusInternalServerError).GetGRPC()
		}
	}
	return resp.(*pb.GetUserResponse), nil
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
