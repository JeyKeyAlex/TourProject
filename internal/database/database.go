package database

import (
	"github.com/JeyKeyAlex/TourProject/internal/entities"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RWDBOperationer interface {
	GetUserList() ([]entities.User, error)
}
type dbp struct {
	db *pgxpool.Pool
}

type RWDBOperation dbp

// NewDBP creates a new database instance
func NewDBP(pool *pgxpool.Pool) *RWDBOperation {
	return &RWDBOperation{db: pool}
}
