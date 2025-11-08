package user

import (
	"context"
	"github.com/JeyKeyAlex/TourProject/internal/database"
	"github.com/JeyKeyAlex/TourProject/internal/entities"
	"github.com/rs/zerolog"
)

type IService interface {
	GetUserList(ctx context.Context) (*entities.GetUserListResponse, error)
	GetLogger() *zerolog.Logger
}
type Service struct {
	rwdbOperation database.RWDBOperationer
	logger        *zerolog.Logger
}

func (s *Service) GetLogger() *zerolog.Logger {
	return s.logger
}

func NewService(rwdbOperation database.RWDBOperationer, logger *zerolog.Logger) IService {
	return &Service{
		rwdbOperation: rwdbOperation,
		logger:        logger,
	}
}
