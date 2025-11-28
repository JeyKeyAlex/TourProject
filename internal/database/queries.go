package database

const (
	queryGetUserList  = `SELECT * FROM users.list;`
	queryCreateUser   = `INSERT INTO users.list (name, last_name, middle_name, nickname, email, phone_number) VALUES ($1, $2, $3, $4, $5, $6 ) RETURNING id;`
	queryCGetUserById = `SELECT email, created_at FROM users.list WHERE id = $1;`
)
