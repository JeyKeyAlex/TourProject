package user

import (
	"context"
	"github.com/JeyKeyAlex/TourProject/internal/config"
	"github.com/JeyKeyAlex/TourProject/internal/database"
	"github.com/JeyKeyAlex/TourProject/internal/entities"
	"github.com/rs/zerolog"
)

type IService interface {
	GetUserList(ctx context.Context) (*entities.GetUserListResponse, error)
	CreateUser(ctx context.Context, req *entities.CreateUserRequest) (*int64, error)

	GetLogger() *zerolog.Logger
}
type Service struct {
	rwdbOperation database.RWDBOperationer
	logger        *zerolog.Logger
	appConfig     *config.Configuration
}

func (s *Service) GetLogger() *zerolog.Logger {
	return s.logger
}

func NewService(rwdbOperation database.RWDBOperationer, logger *zerolog.Logger, appConfig *config.Configuration) IService {
	return &Service{
		rwdbOperation: rwdbOperation,
		logger:        logger,
		appConfig:     appConfig,
	}
}
