package database

import (
	"github.com/JeyKeyAlex/TourProject/internal/entities"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DBP struct {
	db *pgxpool.Pool
}

// NewDBP creates a new database instance
func NewDBP(pool *pgxpool.Pool) *DBP {
	return &DBP{db: pool}
}

type RWDBOperationer interface {
	GetUserList() ([]entities.User, error)
}
