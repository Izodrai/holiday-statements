package users

import (
// 	tmpl "../../templates"
// 	"../../tools"
	"github.com/abbot/go-http-auth"
	"net/http"
// 	"../../db"
// 	locAuth "../../auth"
)

func Update(w http.ResponseWriter, r *auth.AuthenticatedRequest) {

// 	nav := tools.GenerateNav(r.Username)
//
// 	login := r.PostFormValue("login")
// 	email := r.PostFormValue("email")
// 	pass := r.PostFormValue("pass")
// 	rights := r.PostFormValue("rights")
// 	addThisUser := r.PostFormValue("updateThisUser")
//
// 	var newUser tools.User
//
// 	info := struct {
// 		Title string
// 		Nav   []string
// 		User  tools.User
// 		Updated bool
// 		Error bool
// 	}{
// 		Title: "updateUser",
// 		Nav: nav,
// 		User : newUser,
// 		Updated : false,
// 		Error : false,
// 	}
//
// 	if addThisUser == "" {
// 		tmpl.TemplateMe(w, r, "admin/users/update.html", info)
// 		return
// 	}
//
// 	if ok := tools.CheckUser(&newUser, login, email, pass, rights, addThisUser); !ok {
// 		info.Error = true
// 		tmpl.TemplateMe(w, r, "admin/users/update.html", info)
// 		return
// 	}
//
// 	if err := db.AddUser(newUser); err != nil {
// 		info.Error = true
// 		tmpl.TemplateMe(w, r, "admin/users/update.html", info)
// 		return
// 	}
//
// 	locAuth.UpdateSecret(newUser)
//
// 	info.User = newUser
// 	info.Added = true
//
// 	tmpl.TemplateMe(w, r, "admin/users/add.html", info)
}
