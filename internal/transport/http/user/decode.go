package user

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"

	"github.com/JeyKeyAlex/TourProject/internal/entities"

	pb "github.com/JeyKeyAlex/TestProject-genproto/user"
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
	req := pb.IdMessage{}

	id := chi.URLParam(r, "id")
	if id == "" {
		return nil, errors.New("missing user id")
	}

	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, errors.New("failed to parse user_id")
	}

	req.Id = userId

	return req, nil
}

func decodeApproveUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var email string

	if email = chi.URLParam(r, "email"); email == "" {
		return nil, errors.New("missing email")
	}

	return email, nil
}

func decodeDeleteUserByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var userId string

	if userId = chi.URLParam(r, "id"); userId == "" {
		return nil, errors.New("missing user id")
	}

	return userId, nil
}
