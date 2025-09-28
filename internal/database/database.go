package database

import (
	"github.com/JeyKeyAlex/TourProject/internal/entities"
	"github.com/jackc/pgx/v5/pgxpool"
)

type dbp struct {
	db *pgxpool.Pool
}

type RWDBOperationer interface {
	GetUserList() ([]entities.User, error)
}
