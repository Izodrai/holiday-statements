package users

import (
	tmpl "../../templates"
	"../../tools"
	"github.com/abbot/go-http-auth"
	"net/http"
	
// 	"../../tools"
// 	"fmt"
// 	"github.com/abbot/go-http-auth"
// 	"golang.org/x/crypto/bcrypt"
// 	"net/http"
)

func HandleUsers(w http.ResponseWriter, r *auth.AuthenticatedRequest) {

	
	nav := tools.GenerateNav(r.Username)
	
	info := struct {
		Title        string
		Nav          []string
		Participants []string
	}{
		Title: "index",
		Nav: nav,
		Participants: []string{
			"Valentin",
			"Emma",
			"Justine",
		},
	}

	tmpl.TemplateMe(w, r, "lib/templates/admin/users/list.html", info)
	
// 	tools.Info("users")
// 	b, _ := bcrypt.GenerateFromPassword([]byte("vp"), 10)
// 	tools.Info(string(b))
// 
// 	fmt.Fprint(w, `<html>`+/*tools.Menu+*/`<p>users</p>`, r.Username)
}
