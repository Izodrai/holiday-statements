package user

import (
	"net/http"
	"../../authentification"
	"../../tools"
	"github.com/gin-gonic/gin"
)

/****
* http://localhost:8080/user/get
* curl -i -X POST -d '{"user_id":<1>, "token":"<token>", "data": <3 or "name"> }' http://localhost:8080/user/get
****/

func Get_user_by_id_or_name(c *gin.Context) {
	
	var json tools.Request
	
	if !authentification.Check_token(c, &json) {
		return
	}
	
	var u tools.User
	
	switch t := json.Data.(type) {
		case string:
			if sU, ok := tools.Users[t]; ok {
				u = sU
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
		
		u.Email = ""
		u.Password = ""
		u.Token = ""
		
		c.JSON(http.StatusOK, gin.H{
			"user": u,
			"code": http.StatusOK,
			"msg": "user found",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusNoContent,
			"msg": "user not found",
		})
		return
	}
}
