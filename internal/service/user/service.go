package user

import (
	"context"
	"github.com/JeyKeyAlex/TourProject/internal/database"
	"github.com/JeyKeyAlex/TourProject/internal/entities"
)

type IService interface {
	GetUserList(ctx context.Context) (*entities.GetUserListResponse, error)
}
type Service struct {
	rwdbOperation database.RWDBOperationer
}

func NewService(rwdbOperation database.RWDBOperationer) IService {
	return &Service{
		rwdbOperation: rwdbOperation,
	}
}
