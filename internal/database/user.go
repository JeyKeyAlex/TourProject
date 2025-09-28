package database

import (
	"context"
	"errors"
	"time"

	"github.com/JeyKeyAlex/TourProject/internal/entities"
)

func (db *DBP) GetUserList() ([]entities.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	rows, err := db.db.Query(ctx, "SELECT * FROM people.data")
	if err != nil {
		err = errors.New("failed to db.GetUserList: " + err.Error())
		return nil, err
	}
	defer rows.Close()

	var users []entities.User

	for rows.Next() {
		var user entities.User
		err = rows.Scan(&user.Id, &user.Email)
		if err != nil {
			err = errors.New("failed to scan in db.GetUserList: " + err.Error())
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		err = errors.New("failed during rows iteration: " + err.Error())
		return nil, err
	}

	return users, nil
}
