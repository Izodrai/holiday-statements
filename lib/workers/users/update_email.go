package users

import (
	"../../authentification"
	"../../tools"
	"github.com/gin-gonic/gin"
	"net/http"
	"../../db"
	"github.com/izodrai/utils/logs"
	"strconv"
)
/****
* http://localhost:8080/users/update/pwd
* curl -i -X POST -d '{"user_id":<1>, "token":"<token>", "data":{"user_to_update":"<1>","new_email":"new_admin"}}' http://localhost:8080/users/update/email
****/

func Update_email(c *gin.Context) {

  var json tools.Request

	if !authentification.Check_token_and_json(c, &json, false) {
		return
	}

	var err error
	var ok bool
	var user_to_update tools.User
  var user_id_to_update int64
  var s_user_id, new_email string

  switch t := json.Data.(type) {
	case map[string]interface{}:
		s_user_id = t["user_to_update"].(string)
		new_email = t["new_email"].(string)
	default:
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "unexpected data value",
		})
		return
	}

	if user_id_to_update, err = strconv.ParseInt(s_user_id, 10, 64); err != nil {
		logs.Error(err)
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "unexpected data value",
		})
	}

	if user_id_to_update == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "bad id for user to update",
		})
		return
	}

	if json.User_id != user_id_to_update {
		if _, ok := tools.Admins_id[json.User_id]; !ok {
			c.JSON(http.StatusForbidden, gin.H{
				"code":  http.StatusForbidden,
				"msg": "invalid rights, you are not allowed to update the password of an another user",
			})
			return
		}
	}

	if user_to_update, ok = tools.Users_id[user_id_to_update]; !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"msg": "this user cannot be updated because he not exist",
		})
		return
	}

	if new_email == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "empty new email",
		})
		return
	}

	user_to_update.Email = new_email

	if err := db.Update_user_email(&user_to_update); err != nil {
		logs.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "Error during email update",
		})
		return
	}

	tools.Users[user_to_update.Login] = user_to_update
	tools.Users_id[user_to_update.Id] = user_to_update

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "email updated",
	})
}
