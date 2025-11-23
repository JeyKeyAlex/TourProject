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
