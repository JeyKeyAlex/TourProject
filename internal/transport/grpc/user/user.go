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

func (s *RPCServer) Create(ctx context.Context, req *pb.CreateRequest) (*emptypb.Empty, error) {
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

func (s *RPCServer) Approve(ctx context.Context, req *pb.ApproveRequest) (*pb.IdMessage, error) {
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

func (s *RPCServer) GetList(ctx context.Context, req *emptypb.Empty) (*pb.GetListResponse, error) {
	_, resp, err := s.getList.ServeGRPC(ctx, req)
	if err != nil {
		var outputError *error_templates.OutputError
		if errors.As(err, &outputError) {
			return nil, outputError.GetGRPC()
		} else {
			return nil, error_templates.New(err.Error(), err, codes.Unknown, http.StatusInternalServerError).GetGRPC()
		}
	}
	return resp.(*pb.GetListResponse), nil
}

func (s *RPCServer) Get(ctx context.Context, req *pb.IdMessage) (*pb.GetResponse, error) {
	_, resp, err := s.get.ServeGRPC(ctx, req)
	if err != nil {
		var outputError *error_templates.OutputError
		if errors.As(err, &outputError) {
			return nil, outputError.GetGRPC()
		} else {
			return nil, error_templates.New(err.Error(), err, codes.Unknown, http.StatusInternalServerError).GetGRPC()
		}
	}
	return resp.(*pb.GetResponse), nil
}
