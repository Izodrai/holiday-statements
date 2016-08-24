package tools

import (
	"time"
	"golang.org/x/crypto/bcrypt"
)
var Users map[string]User
var UsersId map[int64]User
var Admins map[string]User

type User struct {
	Id       int64
	Login    string
	Password string
	Email    string
	Admin    bool
}

type Event struct {
	Id int64
	Reference string
	CreatedAt EventTime
	PromoterId int64
	PromoterName string
}

type EventTime struct {
	TimeStruct time.Time
	TimeString string
	TimeStamp int64
}

func (e *Event) Feed() {
	e.CreatedAt.TimeStruct = time.Unix(0, e.CreatedAt.TimeStamp*int64(time.Second))
	e.CreatedAt.TimeString = e.CreatedAt.TimeStruct.Format("2006-01-02 15:04:05")
	promoter,_ := UsersId[e.PromoterId]
	e.PromoterName = promoter.Login
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
	
	if pass == "" || len([]byte(pass)) < miniLenPwd {
		return false
	}
	
	if rights != "user" && rights != "admin" {
		return false
	}
	
	if _, ok := Users[login]; ok {
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
		"events",
	}
	
	if _, ok := Admins[login]; ok {
		nav = append(nav, "users")
	}
	
	return nav
}