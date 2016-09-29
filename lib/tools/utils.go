package tools

import (
	"fmt"
	"time"
	"crypto/sha256"
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

type Request struct {
	User_id  int64 `form:"user_id" binding:"required"`
	Token    string `form:"token" binding:"required"`
	Data     interface{} `form:"data" binding:"required"`
}


type Login_form struct {
	Login  string `form:"login" binding:"required"`
	Pwd    string `form:"pwd" binding:"required"`
}

func Crypt_sha256(to_hash string) string {
	
	s := sha256.Sum256([]byte(to_hash))
	
	return fmt.Sprintf("%x", s)
}