package tools


var Users map[string]User
var Admins map[string]User


type User struct {
	Login string `json:"Login"`
	Password string `json:"Password"`
	Email string `json:"Email"`
	Admin bool `json:"Admin"`
}

var Menu = `<a href="index">index</a><br><a href="events">events</a><br><a href="users">users</a>`