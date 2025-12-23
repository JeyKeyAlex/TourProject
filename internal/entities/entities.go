package entities

import "time"

type GetUserListResponse struct {
	Count int64  `json:"count"`
	Users []User `json:"users"`
}
type User struct {
	Id          int64      `db:"id" json:"id,omitempty"`
	Name        string     `db:"name" json:"name,omitempty"`
	LastName    string     `db:"last_name" json:"last_name,omitempty"`
	MiddleName  *string    `db:"middle_name" json:"middle_name,omitempty"`
	Nickname    *string    `db:"nickname" json:"nickname,omitempty"`
	Email       string     `db:"email" json:"email,omitempty"`
	PhoneNumber *string    `db:"phone_number" json:"phone_number,omitempty"`
	CreatedAt   *time.Time `db:"created_at" json:"created_at,omitempty"`
}

type CreateUserRequest struct {
	Name        string  `db:"name" json:"name,omitempty"`
	LastName    string  `db:"last_name" json:"last_name,omitempty"`
	MiddleName  *string `db:"middle_name" json:"middle_name,omitempty"`
	Nickname    *string `db:"nickname" json:"nickname,omitempty"`
	Email       string  `db:"email" json:"email,omitempty"`
	PhoneNumber *string `db:"phone_number" json:"phone_number,omitempty"`
}
