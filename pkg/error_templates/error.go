package error_templates

import (
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"net/http"
)

type OutputError struct {
	errorMessage   string
	errorDetail    error
	grpcStatusCode codes.Code
	httpStatusCode int
}

func (e *OutputError) GetHTTP() (int, string) {
	return e.httpStatusCode, e.errorMessage
}

func (e *OutputError) ErrorDetail() error {
	return e.errorDetail
}

func (e *OutputError) Error() string {
	return e.errorMessage
}

// New is the constructor for OutputError/
func New(errorMessage string, errorDetail error, grpcCode codes.Code, httpCode int) *OutputError {
	return &OutputError{
		errorMessage:   errorMessage,
		errorDetail:    errorDetail,
		grpcStatusCode: grpcCode,
		httpStatusCode: httpCode,
	}
}

// WrapErrorDetail is the function that wraps OutputError.errorDetail field by input message/
func WrapErrorDetail(err error, wrapInfo string) error {
	if outErr, ok := err.(*OutputError); ok {
		outErr.errorDetail = errors.Wrap(outErr.errorDetail, wrapInfo)
		return outErr
	}

	return New(wrapInfo, err, codes.Unknown, http.StatusInternalServerError)
}

func ErrorDetailFromError(err error) error {
	if outErr, ok := err.(*OutputError); ok {
		return outErr.ErrorDetail()
	}

	return err
}
