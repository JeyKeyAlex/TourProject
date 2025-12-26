package user

import (
	"github.com/JeyKeyAlex/TourProject/internal/endpoint/user"
	"github.com/JeyKeyAlex/TourProject/internal/transport/grpc/common"

	kitgrpc "github.com/go-kit/kit/transport/grpc"

	pb "github.com/JeyKeyAlex/TestProject-genproto/user"
)

type RPCServer struct {
	create      kitgrpc.Handler
	approve     kitgrpc.Handler
	update      kitgrpc.Handler
	delete      kitgrpc.Handler
	getUser     kitgrpc.Handler
	getUserList kitgrpc.Handler

	pb.UnimplementedUserServiceServer
}

func NewServer(endpoints user.Endpoints, serverOptions []kitgrpc.ServerOption) pb.UserServiceServer {
	return &RPCServer{
		create:      kitgrpc.NewServer(endpoints.Create, common.DecodeRequest, common.EncodeResponse, serverOptions...),
		approve:     kitgrpc.NewServer(endpoints.Approve, common.DecodeRequest, common.EncodeResponse, serverOptions...),
		update:      kitgrpc.NewServer(endpoints.Update, common.DecodeRequest, common.EncodeResponse, serverOptions...),
		delete:      kitgrpc.NewServer(endpoints.Delete, common.DecodeRequest, common.EncodeResponse, serverOptions...),
		getUser:     kitgrpc.NewServer(endpoints.GetUser, common.DecodeRequest, common.EncodeResponse, serverOptions...),
		getUserList: kitgrpc.NewServer(endpoints.GetUserList, common.DecodeRequest, common.EncodeResponse, serverOptions...),
	}
}
