package users

import (
	"github.com/gin-gonic/gin"
)

// http://localhost:8080/users/test

func get_user_by_name(c *gin.Context) {
	
	search_user_name := c.Param("user_name")
	
	id := 10
	user_name := search_user_name
	
	c.JSON(200, gin.H{
		"id": id,
		"name": user_name,
	})
}