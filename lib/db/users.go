package db

import (
	"../tools"
	"database/sql"
)

func Load_users(users map[string]tools.User) error {

	var err error
	var rows *sql.Rows

	rows, err = Db_connect.Query(`select id, login, pwd, email, rights from users`)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {

		var user tools.User

		err = rows.Scan(&user.Id, &user.Login, &user.Password, &user.Email, &user.Admin)
		if err != nil {
			return err
		}
		users[user.Login] = user
	}
	return nil
}