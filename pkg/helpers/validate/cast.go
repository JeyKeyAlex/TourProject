package validate

import (
	"errors"
	"fmt"
	"net/http"

	"google.golang.org/grpc/codes"

	"github.com/JeyKeyAlex/TourProject/pkg/error_templates"
	pkgerr "github.com/JeyKeyAlex/TourProject/pkg/errors"
)

func CastRequest[T any](request any) (T, error) {
	req, ok := request.(T)
	if !ok {
		err := error_templates.New("invalid request fields", errors.New(pkgerr.FailedCastRequest), codes.InvalidArgument, http.StatusBadRequest)
		return *new(T), error_templates.WrapErrorDetail(err, fmt.Sprintf("cannot cast request to %T", req))
	}
	return req, nil
}
