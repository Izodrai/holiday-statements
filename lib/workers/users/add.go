package users

import (
	"../../authentification"
	"../../tools"
	"github.com/gin-gonic/gin"
	"net/http"
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

  var newUser tools.User
  switch t := json.Data.(type) {
  case map[string]interface{}:
      newUser.Login = t["login"].(string)
      newUser.Password = t["password"].(string)
      newUser.Email = t["email"].(string)
    default:
      c.JSON(http.StatusOK, gin.H{
        "code": http.StatusBadRequest,
        "msg":  "unexpected data value",
      })
      return
  }

  c.JSON(http.StatusOK, gin.H{
    "code": http.StatusOK,
    "msg":  "user added",
  })
}
