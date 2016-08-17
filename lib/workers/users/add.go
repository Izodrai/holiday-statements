package users

import (
	tmpl "../../templates"
	"../../tools"
	"github.com/abbot/go-http-auth"
	"net/http"
	"../../db"
)

func Add(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
	
	nav := tools.GenerateNav(r.Username)
	
	login := r.PostFormValue("login")
	email := r.PostFormValue("email")
	pass := r.PostFormValue("pass")
	rights := r.PostFormValue("rights")
	addThisUser := r.PostFormValue("addThisUser")
	
	info := struct {
		Title string
		Nav   []string
		User  tools.User
		Added bool
		Error bool
	}{
		Title: "addUser",
		Nav: nav,
		User : newUser,
		Added : false,
		Error : false,
	}
	
	var newUser tools.User
	
	if ok := CheckUser(&newUser, login, email, pass, rights, addThisUser); !ok {
		info.Error = true
		tmpl.TemplateMe(w, r, "lib/templates/admin/users/add.html", info)
		return
	}
	
	if ok := db.AddUser(&newUser); !ok {
		info.Error = true
		tmpl.TemplateMe(w, r, "lib/templates/admin/users/add.html", info)
		return
	}
	
	info.User = newUser
	info.Added = true
	

	tmpl.TemplateMe(w, r, "lib/templates/admin/users/add.html", info)
}

const miniLenPwd = 4

func CheckUser(newUser *tools.User, login, email, pass, rights, addThisUser string) bool {
	
	if addThisUser != "ok" {
		return false
	}
	
	if login == "" {
		return false
	}
	
	if email == "" {
		return false
	}
	
	if pass == "" && len([]byte(pass)) < miniLenPwd {
		return false
	}
	
	if rights != "user" || rights != "admin" {
		return false
	}
	
	newUser.Login = login
	newUser.Email = email
	newUser.Password = pass
	
	if rights == "admin" {
		newUser.Admin = true
	} else {
		newUser.Admin = false
	}
	return true
}