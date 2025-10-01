package error_templates

import "google.golang.org/grpc/codes"

type OutputError struct {
	errorMessage   string
	errorDetail    error
	grpcStatusCode codes.Code
	httpStatusCode int
}

func (e *OutputError) GetHTTP() (int, string) {
	return e.httpStatusCode, e.errorMessage
}
