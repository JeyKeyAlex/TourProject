package database

const (
	queryGetUserList  = `SELECT id, email FROM users.list;`
	queryCreateUser   = `INSERT INTO users.list (email) VALUES ($1) RETURNING id;`
	queryCGetUserById = `SELECT email, created_at FROM users.list WHERE id = $1;`
)
