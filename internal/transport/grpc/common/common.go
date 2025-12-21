package common

import (
	"context"
	"errors"

	pb "github.com/JeyKeyAlex/TourProject-proto/go-genproto/user"
)

func DecodeRequest(_ context.Context, request interface{}) (interface{}, error) {
	return request, nil
}

// DecodeIdRequest extracts id from pb.Id protobuf message
func DecodeIdRequest(_ context.Context, request interface{}) (interface{}, error) {
	req, ok := request.(*pb.Id)
	if !ok {
		return nil, errors.New("request must be *pb.Id")
	}
	return req.Id, nil
}

func EncodeResponse(_ context.Context, response interface{}) (interface{}, error) {
	return response, nil
}
