package user

import (
	"context"

	"github.com/JeyKeyAlex/TourProject/internal/entities"
	"github.com/JeyKeyAlex/TourProject/pkg/helpers/saga"
)

func (s *Service) CreateUser(ctx context.Context, req *entities.CreateUserRequest) error {
	logger := s.logger.With().Str("service", "Create").Logger()

	err := s.redisDB.SaveUser(ctx, logger, req, s.appConfig)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) ApproveUser(ctx context.Context, email string) (*int64, error) {
	logger := s.logger.With().Str("service", "Approve").Logger()
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
		return s.rwdbOperation.DeleteUserById(ctx, logger, *id)
	})

	err = s.redisDB.DeleteUser(ctx, logger, email, s.appConfig)
	if err != nil {
		saga.Rollback()
		return nil, err
	}

	return id, nil
}

func (s *Service) UpdateUser(ctx context.Context, req *entities.UpdateUserRequest) (*int64, error) {
	logger := s.logger.With().Str("service", "UpdateUser").Logger()

	id, err := s.rwdbOperation.UpdateUser(ctx, logger, req)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (s *Service) DeleteUserById(ctx context.Context, userId int64) error {
	logger := s.logger.With().Str("service", "Delete").Logger()

	err := s.rwdbOperation.DeleteUserById(ctx, logger, userId)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetUserById(ctx context.Context, userId int64) (*entities.User, error) {
	logger := s.logger.With().Str("service", "GetUser").Logger()

	user, err := s.rwdbOperation.GetUserById(ctx, logger, userId)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) GetUserList(ctx context.Context) (*entities.GetUserListResponse, error) {
	logger := s.logger.With().Str("service", "GetUserList").Logger()

	list, err := s.rwdbOperation.GetUserList(ctx, logger)
	if err != nil {
		return nil, err
	}

	return list, nil
}
