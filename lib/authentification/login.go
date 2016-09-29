package authentification

import (
	"time"
	"../tools"
	"net/http"
	"github.com/gin-gonic/gin"
)

// http://localhost:8080/login/

// curl -i -X POST -d '{"login":"test","pwd":"pwd"}' http://localhost:8080/authentification/

func Login(c *gin.Context) {

	var json tools.Login_form
	
	if c.BindJSON(&json) == nil {
		if u, ok := tools.Users[json.Login]; ok {
			if tools.Crypt_sha256(json.Pwd) == u.Password {
				
				u.Last_connection = time.Now()
				u.Token =  tools.Crypt_sha256(u.Login+"-"+u.Last_connection.Format("2006-01-02 15:04:05"))
				
				tools.Connected_users[u.Id] = u
				
				c.JSON(http.StatusOK, gin.H{
					"user_id": u.Id,
					"user_name": u.Login,
					"token": u.Token,
					"msg": "you are logged in",
				})
				return
			}
			
			c.JSON(http.StatusForbidden, gin.H{
				"msg": "the login or the pwd are invalid",
			})
			return
		}
		
		c.JSON(http.StatusForbidden, gin.H{
			"msg": "the login or the pwd are invalid",
		})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"msg": "no login or pwd found in params of the query",
	})
}