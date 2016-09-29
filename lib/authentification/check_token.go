package authentification

import (
	"net/http"
	"../tools"
	"github.com/gin-gonic/gin"
)
/*
 {
	"user_id": 1,
	"token": "",
	"data": {}
}
*/

func Check_token(c *gin.Context) bool {
	var json tools.Request
	
	if c.BindJSON(&json) == nil {
		if u, ok := tools.Connected_users[json.User_id]; ok {
			if json.Token == u.Token {
				return true
			}
		}
	}
	c.JSON(http.StatusForbidden, gin.H{
		"msg": "invalid token",
	})
	return false
}