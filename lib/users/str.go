package users

import (
	"time"
)

var Users map[string]User
var Users_id map[int64]User
var Admins map[string]User

var Connected_users map[int64]User

type User struct {
	Id       int64
	Login    string
	Password string
	Email    string
	Admin    bool
	Friends  []int64
	Token    string
	Last_connection time.Time
}