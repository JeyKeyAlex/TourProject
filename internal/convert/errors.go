package convert

import (
	"errors"
	"net/http"

	"google.golang.org/grpc/codes"

	"github.com/JeyKeyAlex/TourProject/pkg/error_templates"
)

var EmptyBodyErr = error_templates.New(
	"request body must be non-empty",
	errors.New("request body must be non-empty"),
	codes.InvalidArgument, http.StatusBadRequest)
