package postgreSql

import (
	"context"
	"errors"
	"github.com/JeyKeyAlex/TourProject/internal/entities"
	"github.com/rs/zerolog"
)

func (db *RWDBOperation) GetUserList(ctx context.Context, logger zerolog.Logger) (*entities.GetUserListResponse, error) {
	timeout, cancel := context.WithTimeout(ctx, db.config.MaxIdleConnectionTimeout)
	defer cancel()

	rows, err := db.db.Query(timeout, queryGetUserList)
	if err != nil {
		err = errors.New("failed to db.GetUserList: " + err.Error())
		logger.Error().Err(err).Msg("failed to GetUserList")
		return nil, err
	}
	defer rows.Close()

	var users []entities.User

	for rows.Next() {
		var user entities.User
		err = rows.Scan(
			&user.Id,
			&user.Name,
			&user.LastName,
			&user.MiddleName,
			&user.Nickname,
			&user.Email,
			&user.PhoneNumber,
			&user.CreatedAt,
		)
		if err != nil {
			err = errors.New("failed to scan in db.GetUserList: " + err.Error())
			logger.Error().Err(err).Msg("failed to GetUserList")
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		err = errors.New("failed during rows iteration: " + err.Error())
		logger.Error().Err(err).Msg("failed to GetUserList")
		return nil, err
	}

	count := int64(len(users))

	resp := &entities.GetUserListResponse{
		Count: count,
		Users: users,
	}

	return resp, nil
}

func (db *RWDBOperation) ApproveUser(ctx context.Context, logger zerolog.Logger, req *entities.CreateUserRequest) (*int64, error) {
	timeout, cancel := context.WithTimeout(ctx, db.config.MaxIdleConnectionTimeout)
	defer cancel()

	var id int64

	err := db.db.QueryRow(timeout, queryCreateUser,
		req.Name,
		req.LastName,
		req.MiddleName,
		req.Nickname,
		req.Email,
		req.PhoneNumber,
	).Scan(&id)
	if err != nil {
		err = errors.New("failed to scan in ApproveUser: " + err.Error())
		logger.Error().Err(err).Msg("failed to ApproveUser")
		return nil, err
	}

	return &id, nil
}

func (db *RWDBOperation) GetUserById(ctx context.Context, logger zerolog.Logger, userId string) (*entities.User, error) {
	timeout, cancel := context.WithTimeout(ctx, db.config.MaxIdleConnectionTimeout)
	defer cancel()

	var user entities.User

	err := db.db.QueryRow(timeout, queryCGetUserById, userId).Scan(
		&user.Name,
		&user.LastName,
		&user.MiddleName,
		&user.Nickname,
		&user.Email,
		&user.PhoneNumber,
		&user.CreatedAt,
	)
	if err != nil {
		err = errors.New("failed to scan in GetUserById: " + err.Error())
		logger.Error().Err(err).Msg("failed to GetUserById")
		return nil, err
	}

	return &user, nil
}
