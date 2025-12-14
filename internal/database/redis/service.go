package redis

import (
	"context"
	"github.com/JeyKeyAlex/TourProject/internal/config"
	"github.com/JeyKeyAlex/TourProject/internal/entities"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

type redisP struct {
	dbRedis *redis.Client
}

type Redis interface {
	SaveUser(ctx context.Context, logger zerolog.Logger, req *entities.CreateUserRequest, cfg *config.Configuration) error
	GetTempUser(ctx context.Context, logger zerolog.Logger, email string, cfg *config.Configuration) (*entities.CreateUserRequest, error)
	DeleteUser(ctx context.Context, logger zerolog.Logger, email string, cfg *config.Configuration) error
}

func New(connClient *redis.Client) (Redis, error) {
	r := redisP{dbRedis: connClient}
	return &r, nil
}
