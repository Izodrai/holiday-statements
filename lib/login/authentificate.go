package login

import (
	"time"
	"../tools"
	"net/http"
	"../users"
	"github.com/gin-gonic/gin"
)

// http://localhost:8080/login/

// curl -i -X POST -d '{"login":"C-3PO","pwd":"R2D2"}' http://localhost:8080/login/

func authentificate(c *gin.Context) {

	var json login_form
	
	if c.BindJSON(&json) == nil {
		if u, ok := users.Users[json.Login]; ok {
			
			if tools.Crypt_sha256(json.Pwd) == u.Password {
				
				u.Last_connection = time.Now()
				u.Token =  tools.Crypt_sha256(u.Login+"-"+u.Last_connection.Format("2006-01-02 15:04:05"))
				
				users.Connected_users[u.Id] = u
				
				c.JSON(http.StatusOK, gin.H{
					"user_id": u.Id,
					"token": u.Token,
					"msg": "you are logged in",
				})
				return
			}
			
			c.JSON(http.StatusForbidden, gin.H{
				"msg": "bad pwd",
			})
			return
		}
		
		c.JSON(http.StatusForbidden, gin.H{
			"msg": "this login not exist",
		})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"msg": "no login or pwd found in params of the query",
	})
}

type login_form struct {
    Login  string `form:"login" binding:"required"`
    Pwd    string `form:"pwd" binding:"required"`
}


// func Check_authentificate(c *gin.Context) bool {
// 	
// 	if token, ok := users.Connected_users[id]; ok {
// 		if token_to_validate == token {
// 			return true
// 		}
// 	}
// 	return false
// }