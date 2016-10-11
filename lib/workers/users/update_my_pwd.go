package users

import (
	"../../authentification"
	"../../tools"
	"github.com/gin-gonic/gin"
	"net/http"
	"../../db"
	"github.com/izodrai/utils/logs"
)
/****
* http://localhost:8080/users/update/pwd
* curl -i -X
POST -d '{"user_id":<1>, "token":"<token>", "data":{"user_to_update":<1>,"old_password":"admin","new_password":"new_admin"}}' http://localhost:8080/users/update/pwd
****/

func Update_password(c *gin.Context) {

  var json tools.Request

	if !authentification.Check_token_and_json(c, &json, false) {
		return
	}

  var user_to_update int
  var old_password, new_password string

  switch t := json.Data.(type) {
	case map[string]interface{}:
		user_to_update = t["user_to_update"].(int)
		old_password = t["old_password"].(string)
		new_password = t["new_password"].(string)
	default:
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "unexpected data value",
		})
		return
	}

}
