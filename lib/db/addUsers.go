package db

import (
// 	"database/sql"
)

func AddUser() error {
	
	var err error
	var rows *sql.Rows
	
	rows, err = DbConnect.Query("select login, email, rights from users")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var user tool.User
		rows.Scan(&user.Login, &user.Email, &user.Admin)
	}

	return nil
}
