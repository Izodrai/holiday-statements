package authentification

import (
// 	"time"
	"strconv"
	"../tools"
	"github.com/gin-gonic/gin"
	"net/http"
	"math/rand"
)

/****
* http://localhost:8080/login/login
* curl -i -X POST -d '{"login":"<name>","pwd":"<pwd>"}' http://localhost:8080/authentification/login
****/

func Login(c *gin.Context) {

	var json tools.Login_form

	if c.BindJSON(&json) == nil {
		if u, ok := tools.Users[json.Login]; ok {
			if tools.Crypt_sha256(json.Pwd) == u.Password {
				
				u.Token = tools.Crypt_sha256(u.Login + "-" + strconv.Itoa(rand.Int()))
				
				u.Update_activity()
				
				c.JSON(http.StatusOK, gin.H{
					"user_id":   u.Id,
					"user_name": u.Login,
					"token":     u.Token,
					"code":  http.StatusOK,
					"msg":       "you are logged in",
				})
				return
			}

			c.JSON(http.StatusForbidden, gin.H{
				"code":  http.StatusForbidden,
				"msg": "the login or the pwd are invalid",
			})
			return
		}

		c.JSON(http.StatusForbidden, gin.H{
			"code":  http.StatusForbidden,
			"msg": "the login or the pwd are invalid",
		})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"code":  http.StatusBadRequest,
		"msg": "no login or pwd found in params of the query",
	})
}