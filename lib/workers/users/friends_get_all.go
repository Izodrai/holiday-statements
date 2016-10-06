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

	var ok bool
	var us []tools.User
	var json tools.Request
	var my_friends_ids []int64

	if !authentification.Check_token_and_json(c, &json, true) {
		return
	}

	if my_friends_ids, ok = tools.Friends[json.User_id]; !ok {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusNoContent,
			"msg":  "friends not found 1",
		})
		return
	}
	
	for _, friend_id := range my_friends_ids {
		if my_friend, ok := tools.Users_id[friend_id]; ok {
			my_friend.Clean_for_send()
			us = append(us, my_friend)
		}
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
			"msg":  "friends not found 2",
		})
		return
	}
}
