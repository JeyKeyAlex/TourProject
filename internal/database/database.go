package database

import (
	"github.com/JeyKeyAlex/TourProject/internal/entities"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RWDBOperationer interface {
	GetUserList() (*entities.GetUserListResponse, error)
}
type dbp struct {
	db *pgxpool.Pool
}

type RWDBOperation dbp

// NewRWDBOperationer creates a new database instance
func NewRWDBOperationer(pool *pgxpool.Pool) *RWDBOperation {
	return &RWDBOperation{db: pool}
}
