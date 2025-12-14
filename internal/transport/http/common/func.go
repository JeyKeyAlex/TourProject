package common

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/JeyKeyAlex/TourProject/internal/config"
	"github.com/JeyKeyAlex/TourProject/pkg/error_templates"
)

// EncodeErrorResponse is a function that forms the error response by error message and response code got from codeFrom function.
func EncodeErrorResponse(ctx context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set(config.HeaderContentTypeKey, config.HeaderContentTypeJSON)
	if outputError, ok := (err).(*error_templates.OutputError); ok {
		code, errMessage := outputError.GetHTTP()
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(map[string]interface{}{"error": errMessage})
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error()})
	}
}
