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

func Check_token(c *gin.Context, json *tools.Request) bool {
	
	if c.BindJSON(json) == nil {
		if u, ok := tools.Connected_users[json.User_id]; ok {
			if json.Token == u.Token {
				return true
			}
		}
		c.JSON(http.StatusForbidden, gin.H{
			"msg": "invalid token",
		})
		return false
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"msg": "invalid parameters",
	})
	return false
}