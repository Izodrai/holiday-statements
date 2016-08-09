package tools

var Users map[string]User
var Admins map[string]User

type User struct {
	Login    string
	Password string
	Email    string
	Admin    bool
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