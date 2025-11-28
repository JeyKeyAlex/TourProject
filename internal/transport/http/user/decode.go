package user

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"net/http"

	"github.com/JeyKeyAlex/TourProject/internal/entities"
)

func decodeEmptyRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	return struct{}{}, nil
}

func decodeCreateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	user := &entities.CreateUserRequest{}

	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		return nil, err
	}

	return user, nil
}

func decodeGetUserByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var userId string

	if userId = chi.URLParam(r, "id"); userId == "" {
		return nil, errors.New("missing user id")
	}

	return userId, nil
}
