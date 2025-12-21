package postgreSql

import (
	"context"

	"github.com/JeyKeyAlex/TourProject/internal/config"
	"github.com/JeyKeyAlex/TourProject/internal/entities"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type RWDBOperationer interface {
	GetUserList(ctx context.Context, logger zerolog.Logger) (*entities.GetUserListResponse, error)
	ApproveUser(ctx context.Context, logger zerolog.Logger, req *entities.CreateUserRequest) (*int64, error)
	GetUserById(ctx context.Context, logger zerolog.Logger, userId string) (*entities.User, error)
	DeleteUserById(ctx context.Context, logger zerolog.Logger, userId string) error
}
type dbp struct {
	db     *pgxpool.Pool
	config *config.DBConfig
}

// New creates a new database instance
func New(pool *pgxpool.Pool, config *config.DBConfig) RWDBOperationer {
	return &dbp{
		db:     pool,
		config: config,
	}
}
