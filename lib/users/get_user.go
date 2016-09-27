package users

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// http://localhost:8080/users/test
// curl -i -X GET http://localhost:8080/users/test

func get_user_by_id_or_name(c *gin.Context) {
	
	search_user_name := c.Param("user_info")
	
	id := 10
	user_name := search_user_name
	
	c.JSON(http.StatusOK, gin.H{
		"id": id,
		"name": user_name,
	})
}