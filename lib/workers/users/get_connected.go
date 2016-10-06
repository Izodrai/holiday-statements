package users

import (
	// 	"fmt"
	// 	"strconv"
	"../../authentification"
	"../../tools"
	"github.com/gin-gonic/gin"
	"net/http"
)

/****
* http://localhost:8080/users/get_connected
* curl -i -X POST -d '{"user_id":<1>, "token":"<token>"}' http://localhost:8080/users/get_connected
****/

func Get_connected(c *gin.Context) {

	var json tools.Request

	if !authentification.Check_token(c, &json, true) {
		return
	}

	var cus []tools.User

	for _, cu := range tools.Connected_users {
		cu.Clean_max_for_send()
		cus = append(cus, cu)
	}

	c.JSON(http.StatusOK, gin.H{
		"users": cus,
		"code":  http.StatusOK,
		"msg":   "users found",
	})
}
