package tools

import (
	"golang.org/x/crypto/bcrypt"
)
var Users map[string]User
var Admins map[string]User

type User struct {
	Login    string
	Password string
	Email    string
	Admin    bool
}

func (u *User) GeneratePwd() error {
	b, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		return err
	}
	u.Password = string(b)
	return nil
}

const miniLenPwd = 4

func CheckUser(newUser *User, login, email, pass, rights, addThisUser string) bool {
	
	if addThisUser != "ok" {
		return false
	}
	
	if login == "" {
		return false
	}
	
	if email == "" {
		return false
	}
	
	if pass == "" && len([]byte(pass)) < miniLenPwd {
		return false
	}
	
	if rights != "user" || rights != "admin" {
		return false
	}
	
	newUser.Login = login
	newUser.Email = email
	newUser.Password = pass
	
	if err := newUser.GeneratePwd(); err != nil {
		Error(err)
		return false
	}
	
	if rights == "admin" {
		newUser.Admin = true
	} else {
		newUser.Admin = false
	}
	return true
}

func GenerateNav(login string) []string {
	nav := []string{
		"index",
		"events",
	}
	
	if _, ok := Admins[login]; ok {
		nav = append(nav, "users")
	}
	
	return nav
}