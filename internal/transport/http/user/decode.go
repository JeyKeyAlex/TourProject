package user

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"

	pb "github.com/JeyKeyAlex/TestProject-genproto/user"
)

func decodeEmptyRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	return struct{}{}, nil
}

func decodeCreateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	user := &pb.CreateUserRequest{}

	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		return nil, err
	}

	return user, nil
}

func decodeApproveUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := &pb.ApproveUserRequest{}

	email := chi.URLParam(r, "email")
	if email == "" {
		return nil, errors.New("missing email")
	}

	req.Email = email

	return req, nil
}

func decodeUpdateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := &pb.UpdateUserRequest{}

	id := chi.URLParam(r, "id")
	if id == "" {
		return nil, errors.New("missing id")
	}

	userId, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.New("failed to parse user_id")
	}

	req.Id = int64(userId)

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return nil, err
	}

	return req, nil
}

func decodeDeleteUserByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := &pb.IdMessage{}

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

func decodeGetUserByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := &pb.IdMessage{}

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
