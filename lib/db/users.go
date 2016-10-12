package db

import (
	"../tools"
	"database/sql"
)

func load_users(users map[string]tools.User) error {
	var err error
	var rows *sql.Rows

	rows, err = Db_connect.Query(`select id, login, pwd, email, rights from users where active = 1`)
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

func Create_user(new_user *tools.User) error {

	res, err := Db_connect.Exec("INSERT INTO users (login, pwd, email, rights, active) VALUES (?,?,?,0,1)", new_user.Login, new_user.Password, new_user.Email)
	if err != nil {
		return err
	}

	if res != nil {
		new_user.Id, err = res.LastInsertId()
		if err != nil {
			return err
		}
	}
	return nil
}

func Update_user_password(user_to_update *tools.User) error {

	_, err := Db_connect.Exec("UPDATE users SET pwd = ? WHERE id = ?", user_to_update.Password, user_to_update.Id)
	if err != nil {
		return err
	}

	return nil
}
