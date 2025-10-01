package user

import (
	"context"
	"github.com/JeyKeyAlex/TourProject/internal/entities"
)

func (s *Service) GetUserList(ctx context.Context) ([]entities.User, error) {
	list, err := s.rwdbOperation.GetUserList()
	if err != nil {
		return nil, err
	}
	return list, nil
}
