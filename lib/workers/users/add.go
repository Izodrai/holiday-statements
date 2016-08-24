package users

import (
	tmpl "../../templates"
	"../../tools"
	"github.com/abbot/go-http-auth"
	"net/http"
	"../../db"
	locAuth "../../auth"
)

func Add(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
	
	nav := tools.GenerateNav(r.Username)
	
	login := r.PostFormValue("login")
	email := r.PostFormValue("email")
	pass := r.PostFormValue("pass")
	rights := r.PostFormValue("rights")
	addThisUser := r.PostFormValue("addThisUser")
	
	var newUser tools.User
	
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
	
	if addThisUser == "" {
		tmpl.TemplateMe(w, r, "lib/templates/admin/users/add.html", info)
		return
	}
	
	if ok := tools.CheckUser(&newUser, login, email, pass, rights, addThisUser); !ok {
		info.Error = true
		tmpl.TemplateMe(w, r, "lib/templates/admin/users/add.html", info)
		return
	}
	
	if err := db.AddUser(&newUser); err != nil {
		info.Error = true
		tmpl.TemplateMe(w, r, "lib/templates/admin/users/add.html", info)
		return
	}
	
	locAuth.UpdateSecret(newUser)
	
	info.User = newUser
	info.Added = true
	
	tmpl.TemplateMe(w, r, "lib/templates/admin/users/add.html", info)
}

