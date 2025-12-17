package user

import (
	"context"
	"strconv"

	"github.com/JeyKeyAlex/TourProject/internal/entities"
	"github.com/JeyKeyAlex/TourProject/pkg/helpers/saga"
)

func (s *Service) GetUserList(ctx context.Context) (*entities.GetUserListResponse, error) {
	logger := s.logger.With().Str("service", "GetUserList").Logger()

	list, err := s.rwdbOperation.GetUserList(ctx, logger)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (s *Service) CreateUser(ctx context.Context, req *entities.CreateUserRequest) error {
	logger := s.logger.With().Str("service", "CreateUser").Logger()

	err := s.redisDB.SaveUser(ctx, logger, req, s.appConfig)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetUserById(ctx context.Context, userId string) (*entities.User, error) {
	logger := s.logger.With().Str("service", "GetUserById").Logger()

	user, err := s.rwdbOperation.GetUserById(ctx, logger, userId)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) DeleteUserById(ctx context.Context, userId string) error {
	logger := s.logger.With().Str("service", "DeleteUserById").Logger()

	err := s.rwdbOperation.DeleteUserById(ctx, logger, userId)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) ApproveUser(ctx context.Context, email string) (*int64, error) {
	logger := s.logger.With().Str("service", "ApproveUser").Logger()
	saga := saga.New()

	user, err := s.redisDB.GetTempUser(ctx, logger, email, s.appConfig)
	if err != nil {
		return nil, err
	}

	id, err := s.rwdbOperation.ApproveUser(ctx, logger, user)
	if err != nil {
		return nil, err
	}

	saga.AddRollbackFunc(func() error {
		userId := strconv.FormatInt(*id, 10)
		return s.rwdbOperation.DeleteUserById(ctx, logger, userId)
	})

	err = s.redisDB.DeleteUser(ctx, logger, email, s.appConfig)
	if err != nil {
		saga.Rollback()
		return nil, err
	}

	return id, nil
}
