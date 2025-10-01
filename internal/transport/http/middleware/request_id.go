package middleware

import (
	"google.golang.org/grpc/metadata"
	"net/http"
)

const RequestIDKey = "X-Request-ID"

func RequestID(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		requestID := r.Header.Get(RequestIDKey)
		md := metadata.Pairs(RequestIDKey, requestID)
		ctx = metadata.NewIncomingContext(ctx, md)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
