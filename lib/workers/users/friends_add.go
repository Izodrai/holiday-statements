package users

import (
	"../../authentification"
	"../../db"
	"../../tools"
	"github.com/gin-gonic/gin"
	"github.com/izodrai/utils/logs"
	"net/http"
	"strconv"
)

/****
* http://localhost:8080/users/get/all
* curl -i -X POST -d '{"user_id":<1>, "token":"<token>", "data":{"friend_to_add":"<1>"}}' http://localhost:8080/users/friends/add
****/

func Friends_add(c *gin.Context) {

	var json tools.Request

	if !authentification.Check_token_and_json(c, &json, false) {
		return
	}

	var err error
	var ok bool
	var my_friends []int64
	var new_friend_id int64
	var snew_friend_id string

	switch t := json.Data.(type) {
	case map[string]interface{}:
		snew_friend_id = t["friend_to_add"].(string)
	default:
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "unexpected data value",
		})
		return
	}

	if new_friend_id, err = strconv.ParseInt(snew_friend_id, 10, 64); err != nil {
		logs.Error(err)
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "unexpected data value",
		})
	}

	if new_friend_id == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "bad id for friend to add",
		})
		return
	}

	if _, ok = tools.Users_id[new_friend_id]; !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "this friend cannot be added because he not exist",
		})
		return
	}

	if my_friends, ok = tools.Friends[json.User_id]; ok {
		for _, friend_id := range my_friends {
			if new_friend_id == friend_id {
				c.JSON(http.StatusBadRequest, gin.H{
					"code": http.StatusBadRequest,
					"msg":  "this user is already your friend",
				})
				return
			}
		}
	}

	if err := db.Add_friend(json.User_id, new_friend_id); err != nil {
		logs.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "Error during friend adding",
		})
		return
	}

	my_friends, _ = tools.Friends[json.User_id]
	my_friends = append(my_friends, new_friend_id)
	tools.Friends[json.User_id] = my_friends

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "friend added",
	})
}
