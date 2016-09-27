package users

import (
	
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
	Friends []int64
}