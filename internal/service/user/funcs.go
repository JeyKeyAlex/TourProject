package user

import (
	"context"
	"github.com/JeyKeyAlex/TourProject/internal/entities"
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
	logger := s.logger.With().Str("service", "ApproveUser").Logger()

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

func (s *Service) ApproveUser(ctx context.Context, email string) (*int64, error) {
	logger := s.logger.With().Str("service", "ApproveUser").Logger()

	user, err := s.redisDB.GetTempUser(ctx, logger, email, s.appConfig)
	if err != nil {
		return nil, err
	}

	id, err := s.rwdbOperation.ApproveUser(ctx, logger, user)
	if err != nil {
		return nil, err
	}
	return id, nil
}
