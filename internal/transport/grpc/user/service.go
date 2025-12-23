package user

import (
	"github.com/JeyKeyAlex/TourProject/internal/endpoint/user"
	"github.com/JeyKeyAlex/TourProject/internal/transport/grpc/common"

	kitgrpc "github.com/go-kit/kit/transport/grpc"

	pb "github.com/JeyKeyAlex/TourProject-proto/go-genproto/user"
)

type RPCServer struct {
	create  kitgrpc.Handler
	approve kitgrpc.Handler
	delete  kitgrpc.Handler
	getList kitgrpc.Handler
	get     kitgrpc.Handler

	pb.UnimplementedUserServiceServer
}

// NewServer is a constructor for creating a new instance of a gRPC server(RPCServer structure).
func NewServer(endpoints user.Endpoints, serverOptions []kitgrpc.ServerOption) pb.UserServiceServer {
	return &RPCServer{
		create:  kitgrpc.NewServer(endpoints.Create, common.DecodeRequest, common.EncodeResponse, serverOptions...),
		approve: kitgrpc.NewServer(endpoints.Approve, common.DecodeRequest, common.EncodeResponse, serverOptions...),
		delete:  kitgrpc.NewServer(endpoints.Delete, common.DecodeRequest, common.EncodeResponse, serverOptions...),
		getList: kitgrpc.NewServer(endpoints.GetUserList, common.DecodeRequest, common.EncodeResponse, serverOptions...),
		get:     kitgrpc.NewServer(endpoints.GetUserById, common.DecodeRequest, common.EncodeResponse, serverOptions...),
	}
}
