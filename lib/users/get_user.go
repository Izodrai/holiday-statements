package users

import (
// 	"../tools"
	"strconv"
	"net/http"
// 	"encoding/json"
	"github.com/gin-gonic/gin"
)

// http://localhost:8080/users/test
// curl -i -X GET http://localhost:8080/users/test

func get_user_by_id_or_name(c *gin.Context) {
	
	var u User
	
	search := c.Param("user_info")
	
	if id, err := strconv.ParseInt(search, 10, 64); err == nil {
		if sU, ok := Users_id[id]; ok {
			u = sU
		}
	} else if sU, ok := Users[search]; ok {
		u = sU
	}
		
	if u.Id != 0 {
		c.JSON(http.StatusOK, gin.H{
			"user": u,
			"msg": "user found",
		})
		return
	} else {
		c.JSON(http.StatusNoContent, gin.H{
			"search": search,
			"msg": "user not found",
		})
		return
	}
}