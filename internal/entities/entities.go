package entities

type User struct {
	Id    int    `db:"id" json:"id,omitempty"`
	Email string `db:"email" json:"email,omitempty"`
}
