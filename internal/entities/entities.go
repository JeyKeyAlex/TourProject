package entities

type GetUserListResponse struct {
	Count int64  `json:"count"`
	Users []User `json:"users"`
}
type User struct {
	Id    int    `db:"id" json:"id,omitempty"`
	Email string `db:"email" json:"email,omitempty"`
}

type CreateUserRequest struct {
	Email string `db:"email" json:"email,omitempty"`
}
