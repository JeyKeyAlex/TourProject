package user

import (
	"context"

	"github.com/JeyKeyAlex/TourProject/internal/config"
	"github.com/JeyKeyAlex/TourProject/internal/database/postgreSql"
	"github.com/JeyKeyAlex/TourProject/internal/database/redis"
	"github.com/JeyKeyAlex/TourProject/internal/entities"

	"buf.build/go/protovalidate"
	"github.com/rs/zerolog"
)

type IService interface {
	GetUserList(ctx context.Context) (*entities.GetUserListResponse, error)
	GetUserById(ctx context.Context, userId int64) (*entities.User, error)

	CreateUser(ctx context.Context, req *entities.CreateUserRequest) error
	ApproveUser(ctx context.Context, email string) (*int64, error)
	UpdateUser(ctx context.Context, req *entities.UpdateUserRequest) (*int64, error)
	DeleteUserById(ctx context.Context, userId int64) error

	GetLogger() *zerolog.Logger
	GetValidator() protovalidate.Validator
}
type Service struct {
	rwdbOperation postgreSql.RWDBOperationer
	redisDB       redis.Redis
	logger        *zerolog.Logger
	appConfig     *config.Configuration
	validator     protovalidate.Validator
}

func (s *Service) GetLogger() *zerolog.Logger {
	return s.logger
}
func (s *Service) GetValidator() protovalidate.Validator {
	return s.validator
}

func NewService(rwdbOperation postgreSql.RWDBOperationer, redisDB redis.Redis, validator protovalidate.Validator, logger *zerolog.Logger, appConfig *config.Configuration) IService {
	return &Service{
		rwdbOperation: rwdbOperation,
		redisDB:       redisDB,
		logger:        logger,
		appConfig:     appConfig,
		validator:     validator,
	}
}
