package users

import (
	"github.com/abbot/go-http-auth"
	"net/http"
	
	wu "../../workers/users"
)

func HandleUsers(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
	
	params := r.URL.Query()
	
	if _, ok := params["add"]; ok {
		wu.Add(w,r)
		return 
	}
	
// 	if _, ok := params["update"]; ok {
// 		wu.Update(w,r)
// 		return 
// 	}
	
	wu.List(w,r)
}
