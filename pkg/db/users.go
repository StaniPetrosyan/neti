package db

import "database/sql"

type User struct {
	Username string
	Password string
}

type Users interface {
	Add(user User) bool
	FindBy(username string) (string, string)
}

type PostgresUsers struct {
	Psql *sql.DB
}

func (u *PostgresUsers) Add(user User) bool {
	insertStmt := `insert into users("username", "password") values($1, $2)`
	_, err := u.Psql.Exec(insertStmt, user.Username, user.Password)

	return err == nil
}

func (u *PostgresUsers) FindBy(username string) (string, string) {
	row := u.Psql.QueryRow(`SELECT * FROM users where username = $1`, username)
	var foundUsername string
	var foundPassword string
	row.Scan(&foundUsername, &foundPassword)

	return foundUsername, foundPassword
}
