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
* http://localhost:8080/users/get/connected
* curl -i -X POST -d '{"user_id":<1>, "token":"<token>"}' http://localhost:8080/users/get/connected
****/

func Get_connected(c *gin.Context) {

	var json tools.Request

	if !authentification.Check_token_and_json(c, &json, true) {
		return
	}

	var us []tools.User

	for _, cu := range tools.Connected_users {
		cu.Clean_for_send()
		us = append(us, cu)
	}

	if len(us) != 0 {
		c.JSON(http.StatusOK, gin.H{
			"users": us,
			"code":  http.StatusOK,
			"msg":   "users found",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusNoContent,
			"msg":  "users not found",
		})
		return
	}
}
