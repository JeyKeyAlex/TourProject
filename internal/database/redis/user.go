package redis

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/JeyKeyAlex/TourProject/internal/config"
	"github.com/JeyKeyAlex/TourProject/internal/entities"
	"github.com/JeyKeyAlex/TourProject/pkg/error_templates"
	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
)

func (r *redisP) SaveUser(ctx context.Context, logger zerolog.Logger, req *entities.CreateUserRequest, cfg *config.Configuration) error {
	timeoutCtx, cancel := context.WithTimeout(ctx, cfg.Redis.Timeout)
	defer cancel()

	const oper = "database.Redis.SaveUser"

	// Convert struct fields to field-value pairs for HSet
	fields := map[string]interface{}{
		"name":      req.Name,
		"last_name": req.LastName,
		"email":     req.Email,
	}

	if req.MiddleName != nil {
		fields["middle_name"] = *req.MiddleName
	}
	if req.Nickname != nil {
		fields["nickname"] = *req.Nickname
	}
	if req.PhoneNumber != nil {
		fields["phone_number"] = *req.PhoneNumber
	}

	cmd := r.dbRedis.HSet(timeoutCtx, req.Email, fields)
	if err := cmd.Err(); err != nil {
		logger.Error().Stack().Err(err).Msg(oper)
		return fmt.Errorf("%s: %w", oper, err)
	}
	return nil
}

func (r *redisP) GetTempUser(ctx context.Context, logger zerolog.Logger, email string, cfg *config.Configuration) (*entities.CreateUserRequest, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, cfg.Redis.Timeout)
	defer cancel()

	var user entities.CreateUserRequest

	resultMap := r.dbRedis.HGetAll(timeoutCtx, email)
	if err := resultMap.Err(); err != nil {
		logger.Error().Stack().Err(err).Msg("failed to get temporary user")
		return nil, error_templates.New(err.Error(), err, codes.Internal, http.StatusInternalServerError)
	}

	userMap := resultMap.Val()
	if len(userMap) == 0 || userMap == nil {
		err := errors.New("no users found, you need to register via CreateUser before approval")
		logger.Error().Stack().Err(err).Msg("failed to get temporary user")
		return nil, error_templates.New(err.Error(), err, codes.NotFound, http.StatusNotFound)
	}

	user.Name = userMap["name"]
	user.LastName = userMap["last_name"]

	middleName := userMap["middle_name"]
	user.MiddleName = &middleName

	nick := userMap["nickname"]
	user.Nickname = &nick

	user.Email = userMap["email"]

	phone := userMap["phone_number"]
	user.PhoneNumber = &phone

	return &user, nil
}

func (r *redisP) DeleteUser(ctx context.Context, logger zerolog.Logger, email string, cfg *config.Configuration) error {
	timeoutCtx, cancel := context.WithTimeout(ctx, cfg.Redis.Timeout)
	defer cancel()

	const oper = "pkg.database.DeleteSession"

	cmd := r.dbRedis.Del(timeoutCtx, email)
	if cmd.Val() == 0 {
		err := errors.New("failed to delete temporary user")
		logger.Error().Stack().Err(err).Msg(oper)
		return error_templates.New(err.Error(), err, codes.InvalidArgument, http.StatusBadRequest)
	}

	return nil
}
