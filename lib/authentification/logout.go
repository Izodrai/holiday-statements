package authentification

import (
	"../tools"
	"github.com/gin-gonic/gin"
	"net/http"
)

/****
* http://localhost:8080/login/login
* curl -i -X POST -d '{"login":"<name>","pwd":"<pwd>"}' http://localhost:8080/authentification/logout
****/

func Logout(c *gin.Context) {
	var json tools.Request

	if !Check_token(c, &json, false) {
		return
	}
	
	u, _ := tools.Connected_users[json.User_id]
	
	u.Update_activity()

	delete(tools.Connected_users, json.User_id)
	
	c.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"msg":   "you have been disconnected",
	})
}
