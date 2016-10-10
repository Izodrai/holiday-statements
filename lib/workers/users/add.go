package users

import (
	"../../authentification"
	"../../tools"
	"github.com/gin-gonic/gin"
	"net/http"
	"../../db"
)

/****
* http://localhost:8080/users/add
* curl -i -X POST -d '{"user_id":<1>, "token":"<token>", "data":{"login":"user4","password":"user4","email":"user4@memail.com"}}' http://localhost:8080/users/add
****/

func Add(c *gin.Context) {

	var json tools.Request

	if !authentification.Check_token_and_json(c, &json, false) {
		return
	}

	var new_user tools.User
	switch t := json.Data.(type) {
	case map[string]interface{}:
		new_user.Login = t["login"].(string)
		new_user.Password = t["password"].(string)
		new_user.Email = t["email"].(string)
	default:
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "unexpected data value",
		})
		return
	}

	if new_user.Login == "" || len(new_user.Login) <= 4 || new_user.Password == "" || len(new_user.Password) <= 4 || new_user.Email == "" || len(new_user.Email) <= 4 {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "empty login / email / password or len < 4",
		})
		return
	}

	if _, ok := tools.Users[new_user.Login]; ok {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "this user already exist",
		})
		return
	}

	new_user.Password = tools.Crypt_sha256(new_user.Password)

	if err := db.Create_user(&new_user); err != nil {
		tools.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"msg":  "Error during user insertion",
		})
		return
	}

	tools.Users[new_user.Login] = new_user
	tools.Users_id[new_user.Id] = new_user

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "user added",
	})
}
