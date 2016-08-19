package db

import (
	"..//tools"
)

func AddUser(newUser tools.User) error {
	_, err := DbConnect.Exec("INSERT INTO users( login, pwd, email, rights) VALUES (?,?,?,?)", newUser.Login, newUser.Password, newUser.Email, newUser.Admin)
	if err != nil {
		return err
	}
	return nil
}