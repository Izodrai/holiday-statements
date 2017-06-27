package users

import (
	tmpl "../../templates"
	"../../tools"
	"github.com/abbot/go-http-auth"
	"net/http"
)

func List(w http.ResponseWriter, r *auth.AuthenticatedRequest) {

	nav := tools.GenerateNav(r.Username)

	var users []tools.User
	var admins []tools.User

	for _, user := range tools.Users {
		users = append(users, user)
	}

	for _, user := range tools.Admins {
		admins = append(admins, user)
	}

	info := struct {
		Title string
		Nav   []string
		Users []tools.User
		Admins []tools.User
	}{
		Title: "userList",
		Nav: nav,
		Users: users,
		Admins: admins,
	}

	tmpl.TemplateMe(w, r, "admin/users/list.html", info)
}
