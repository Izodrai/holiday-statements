package users

import (
	"github.com/gin-gonic/gin"
)


func Handler(router *gin.Engine) {
	
	users := router.Group("/users")
	{
		users.GET("/:user_info", get_user_by_id_or_name)
	}
}