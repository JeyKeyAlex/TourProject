package validate

import (
	"buf.build/go/protovalidate"
	"errors"
	"fmt"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/proto"

	"github.com/JeyKeyAlex/TourProject/pkg/error_templates"
	pkgerr "github.com/JeyKeyAlex/TourProject/pkg/errors"
)

func CastValidateRequest[T proto.Message](validator protovalidate.Validator, request any) (T, error) {
	req, ok := request.(T)
	if !ok {
		err := error_templates.New("invalid request fields", errors.New(pkgerr.FailedCastRequest), codes.InvalidArgument, http.StatusBadRequest)
		return *new(T), error_templates.WrapErrorDetail(err, fmt.Sprintf("cannot cast request to %T", req))
	}

	err := validator.Validate(req)
	if err != nil {
		return req, error_templates.New(err.Error(), err, codes.InvalidArgument, http.StatusBadRequest)
	}

	return req, nil
}
