package user

import (
	"context"
	"github.com/JeyKeyAlex/TourProject/internal/config"
	"github.com/JeyKeyAlex/TourProject/internal/database/postgreSql"
	"github.com/JeyKeyAlex/TourProject/internal/database/redis"
	"github.com/JeyKeyAlex/TourProject/internal/entities"

	"github.com/rs/zerolog"
)

type IService interface {
	GetUserList(ctx context.Context) (*entities.GetUserListResponse, error)
	CreateUser(ctx context.Context, req *entities.CreateUserRequest) error
	ApproveUser(ctx context.Context, email string) (*int64, error)
	GetUserById(ctx context.Context, userId string) (*entities.User, error)
	DeleteUserById(ctx context.Context, userId string) error

	GetLogger() *zerolog.Logger
}
type Service struct {
	rwdbOperation postgreSql.RWDBOperationer
	redisDB       redis.Redis
	logger        *zerolog.Logger
	appConfig     *config.Configuration
}

func (s *Service) GetLogger() *zerolog.Logger {
	return s.logger
}

func NewService(rwdbOperation postgreSql.RWDBOperationer, redisDB redis.Redis, logger *zerolog.Logger, appConfig *config.Configuration) IService {
	return &Service{
		rwdbOperation: rwdbOperation,
		redisDB:       redisDB,
		logger:        logger,
		appConfig:     appConfig,
	}
}
