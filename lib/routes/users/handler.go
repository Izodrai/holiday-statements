package users

import (
	"fmt"
	"net/http"
	"../../tools"
	"golang.org/x/crypto/bcrypt"
	"github.com/abbot/go-http-auth"
)

func HandleUsers(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
	
	tools.Info("users")
	b,_ :=bcrypt.GenerateFromPassword([]byte("vp"),10)
	tools.Info(string(b))
	
	fmt.Fprint(w,`<html>`+tools.Menu+`<p>users</p>`,r.Username)
}