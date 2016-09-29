package handler

import (
	"../workers/users"
	"../authentification"
	"github.com/gin-gonic/gin"
)


func Handler(router *gin.Engine) {
	a := router.Group("/authentification")
	{
		a.POST("/", authentification.Login)
	}
	u := router.Group("/users")
	{
		u.POST("/:user_info", users.Get_user_by_id_or_name)
	}
}