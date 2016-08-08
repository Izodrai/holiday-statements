package tools

var Users map[string]User
var Admins map[string]User

type User struct {
	Login    string
	Password string
	Email    string
	Admin    bool
}
