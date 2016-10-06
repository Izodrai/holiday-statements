package users

import (
	"../../authentification"
	"../../tools"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

/****
* http://localhost:8080/users/get
* curl -i -X POST -d '{"user_id":<1>, "token":"<token>", "data": <3 or "name" or "email"> }' http://localhost:8080/users/get
****/

func Search(c *gin.Context) {

	var json tools.Request

	if !authentification.Check_token_and_json(c, &json, false) {
		return
	}

	var u tools.User

	switch t := json.Data.(type) {
	case string:
		if sU, ok := tools.Users[t]; ok {
			u = sU
		} else {
			for _, sU := range tools.Users {
				if strings.Replace(strings.ToLower(t)," ","",-1) == sU.Email {
					u = sU
					break
				}
			}
		}
	case int:
		if sU, ok := tools.Users_id[int64(t)]; ok {
			u = sU
		}
	case float64:
		if sU, ok := tools.Users_id[int64(t)]; ok {
			u = sU
		}
	}

	if u.Id != 0 {
		u.Clean_max_for_send()

		c.JSON(http.StatusOK, gin.H{
			"user": u,
			"code": http.StatusOK,
			"msg":  "user found",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusNoContent,
			"msg":  "user not found",
		})
		return
	}
}
