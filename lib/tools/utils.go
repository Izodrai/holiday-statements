package tools

import (
	"crypto/sha256"
	"fmt"
	"time"
)

var Users map[string]User
var Users_id map[int64]User
var Admins map[string]User
var Connected_users map[int64]User

type User struct {
	Id              int64
	Login           string
	Password        string
	Email           string
	Admin           bool
	Friends         []int64
	Token           string
	Last_activity time.Time
}

type Request struct {
	User_id int64       `form:"user_id" binding:"required"`
	Token   string      `form:"token" binding:"required"`
	Data    interface{} `form:"data"`
}

type Login_form struct {
	Login string `form:"login" binding:"required"`
	Pwd   string `form:"pwd" binding:"required"`
}

func Crypt_sha256(to_hash string) string {

	s := sha256.Sum256([]byte(to_hash))

	return fmt.Sprintf("%x", s)
}

func (u *User) Update_activity() {
	
	u.Last_activity = time.Now()
	
	if u2, ok := Connected_users[u.Id]; ok {
		u2.Last_activity = u.Last_activity
		Connected_users[u2.Id] = u2
	} else {
		Connected_users[u.Id] = *u
	}
	
	if u2, ok := Users_id[u.Id]; ok {
		u2.Last_activity = u.Last_activity
		Users_id[u2.Id] = u2
	}
	
	if u2, ok := Users[u.Login]; ok {
		u2.Last_activity = u.Last_activity
		Users[u2.Login] = u2
	}
	
	if u2, ok := Admins[u.Login]; ok {
		u2.Last_activity = u.Last_activity
		Admins[u2.Login] = u2
	}
}