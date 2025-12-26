package postgreSql

const (
	queryGetUserList    = `SELECT * FROM users.list;`
	queryGetUserById    = `SELECT id, name, last_name, middle_name, nickname, email, phone_number, created_at FROM users.list WHERE id = $1;`
	queryDeleteUserById = `DELETE FROM users.list WHERE id = $1;`
	queryCreateUser     = `INSERT INTO users.list (name, last_name, middle_name, nickname, email, phone_number) VALUES ($1, $2, $3, $4, $5, $6 ) RETURNING id;`
	queryUpdate         = `UPDATE users.list SET 
						name = COALESCE($1, name), 
						last_name = COALESCE($2, last_name), 
						middle_name = COALESCE($3, middle_name), 
						nickname = COALESCE($4, nickname), 
						email = COALESCE($5, email), 
						phone_number = COALESCE($6, phone_number)
						WHERE id = $7 RETURNING id;`
)
