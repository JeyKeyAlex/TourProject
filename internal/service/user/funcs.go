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

func (s *Service) CreateUser(ctx context.Context, req *entities.CreateUserRequest) (*int64, error) {
	logger := s.logger.With().Str("service", "CreateUser").Logger()

	id, err := s.rwdbOperation.CreateUser(ctx, logger, req)
	if err != nil {
		return nil, err
	}
	return id, nil
}

func (s *Service) GetUserById(ctx context.Context, userId string) (*entities.User, error) {
	logger := s.logger.With().Str("service", "GetUserById").Logger()

	user, err := s.rwdbOperation.GetUserById(ctx, logger, userId)
	if err != nil {
		return nil, err
	}

	return user, nil
}
