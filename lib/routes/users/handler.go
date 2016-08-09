package users

import (
	tmpl "../../templates"
	"../../tools"
	"github.com/abbot/go-http-auth"
	"net/http"
	
// 	"fmt"
// 	"github.com/abbot/go-http-auth"
// 	"golang.org/x/crypto/bcrypt"
// 	"net/http"
)

func HandleUsers(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
	
	params := r.URL.Query()
	
	if _, ok := params["add"]; ok {
		Add(w,r)
		return 
	}
	
	List(w,r)
// 	tools.Info("users")
// 	b, _ := bcrypt.GenerateFromPassword([]byte("vp"), 10)
// 	tools.Info(string(b))
// 
// 	fmt.Fprint(w, `<html>`+/*tools.Menu+*/`<p>users</p>`, r.Username)
}

func Add(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
	
	nav := tools.GenerateNav(r.Username)
	
	login := r.PostFormValue("login")
	email := r.PostFormValue("email")
	pass := r.PostFormValue("pass")
	addThisUser := r.PostFormValue("addThisUser")
	
	tools.Info(login)
	tools.Info(email)
	tools.Info(pass)
	tools.Info(addThisUser)
	
	
	
	
	newUser := tools.User{
		Login    :"newUser",
		Password :"",
		Email    :"",
		Admin    :false,
	}
	
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

	tmpl.TemplateMe(w, r, "lib/templates/admin/users/add.html", info)
}

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

	tmpl.TemplateMe(w, r, "lib/templates/admin/users/list.html", info)
}
