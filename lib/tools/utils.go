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

var Friends map[int64][]int64

type User struct {
	Id              int64		`json:"id,omitempty"`
	Login           string		`json:"login,omitempty"`
	Password        string		`json:"password,omitempty"`
	Email           string		`json:"email,omitempty"`
	Admin           bool		`json:"admin,omitempty"`
	Friends         []int64		`json:"friends,omitempty"`
	Token           string		`json:"token,omitempty"`
	Last_activity   time.Time	`json:"last_activity,omitempty"`
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

func (u *User) Clean_for_send() {
	u.Password = ""
	u.Email = ""
	u.Admin = false
	u.Friends = []int64{}
	u.Token = ""
	u.Last_activity = time.Time{}
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