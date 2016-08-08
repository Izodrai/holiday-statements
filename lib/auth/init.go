package auth

import (
	"../db"
	"../tools"
	"database/sql"
)

func Init() error {
	
	tools.Users = make(map[string]tools.User)
	tools.Admins = make(map[string]tools.User)
	
	if err := loadUsers(tools.Users); err != nil {
		return err
	}
	
	for login, user := range tools.Users {
		if user.Admin {
			tools.Admins[login]=user
		}
	}
	
	return nil
}

func loadUsers(users map[string]tools.User) error {
	
	var err error
	var rows *sql.Rows
	
	rows, err = db.DbConnect.Query("select login, pwd, email, rights from users;")
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
		
		users[user.Login]=user
	}
	
	return nil
}