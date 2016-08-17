package users

import (
// 	tmpl "../../templates"
// 	"../../tools"
	"github.com/abbot/go-http-auth"
	"net/http"
	
	wu "../../workers/users"
	
// 	"fmt"
// 	"github.com/abbot/go-http-auth"
// 	"golang.org/x/crypto/bcrypt"
// 	"net/http"
)

func HandleUsers(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
	
	params := r.URL.Query()
	
	if _, ok := params["add"]; ok {
		wu.Add(w,r)
		return 
	}
	
	wu.List(w,r)
// 	tools.Info("users")
// 	b, _ := bcrypt.GenerateFromPassword([]byte("vp"), 10)
// 	tools.Info(string(b))
// 
// 	fmt.Fprint(w, `<html>`+/*tools.Menu+*/`<p>users</p>`, r.Username)
}
