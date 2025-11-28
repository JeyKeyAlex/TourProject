package database

import (
	"context"
	"github.com/JeyKeyAlex/TourProject/internal/config"
	"github.com/JeyKeyAlex/TourProject/internal/entities"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type RWDBOperationer interface {
	GetUserList(ctx context.Context, logger zerolog.Logger) (*entities.GetUserListResponse, error)
	CreateUser(ctx context.Context, logger zerolog.Logger, req *entities.CreateUserRequest) (*int64, error)
	GetUserById(ctx context.Context, logger zerolog.Logger, userId string) (*entities.User, error)
}
type dbp struct {
	db     *pgxpool.Pool
	config *config.DBConfig
}

type RWDBOperation dbp

// NewRWDBOperationer creates a new database instance
func NewRWDBOperationer(pool *pgxpool.Pool, config *config.DBConfig) *RWDBOperation {
	return &RWDBOperation{
		db:     pool,
		config: config,
	}
}
