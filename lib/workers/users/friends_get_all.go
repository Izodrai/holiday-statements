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
* http://localhost:8080/users/get/all
* curl -i -X POST -d '{"user_id":<1>, "token":"<token>"}' http://localhost:8080/users/get/all
****/

func Friends_get_all(c *gin.Context) {

	var json tools.Request

	if !authentification.Check_token_and_json(c, &json, true) {
		return
	}

	me := tools.Users_id[json.User_id]
	
	var us []tools.User
	
	for _, friend_id := range me.Friends {
		friend := tools.Users_id[friend_id]
		friend.Clean_for_send()
		us = append(us, friend)
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
