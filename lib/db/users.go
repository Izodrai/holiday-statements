package db

import (
	"../tools"
	"database/sql"
)

func LoadUsers(users map[string]tools.User) error {

	var err error
	var rows *sql.Rows

	rows, err = DbConnect.Query("select login, pwd, email, rights from users;")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {

		var user tools.User

		err = rows.Scan(&user.Login, &user.Password, &user.Email, &user.Admin)
		if err != nil {
			return err
		}
		users[user.Login] = user
	}
	return nil
}

func AddUser(newUser tools.User) error {
	_, err := DbConnect.Exec("INSERT INTO users( login, pwd, email, rights) VALUES (?,?,?,?)", newUser.Login, newUser.Password, newUser.Email, newUser.Admin)
	if err != nil {
		return err
	}
	return nil
}