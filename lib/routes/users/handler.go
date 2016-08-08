package users

import (
	"../../tools"
	"fmt"
	"github.com/abbot/go-http-auth"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func HandleUsers(w http.ResponseWriter, r *auth.AuthenticatedRequest) {

	tools.Info("users")
	b, _ := bcrypt.GenerateFromPassword([]byte("vp"), 10)
	tools.Info(string(b))

	fmt.Fprint(w, `<html>`+tools.Menu+`<p>users</p>`, r.Username)
}
