package postgreSql

const (
	queryGetUserList    = `SELECT * FROM users.list;`
	queryCreateUser     = `INSERT INTO users.list (name, last_name, middle_name, nickname, email, phone_number) VALUES ($1, $2, $3, $4, $5, $6 ) RETURNING id;`
	queryGetUserById    = `SELECT name, last_name, middle_name, nickname, email, phone_number, created_at FROM users.list WHERE id = $1;`
	queryDeleteUserById = `DELETE FROM users.list WHERE id = $1;`
)
