package handler

import (
	"../workers/user"
	"../authentification"
	"github.com/gin-gonic/gin"
)


func Handler(router *gin.Engine) {
	a := router.Group("/authentification")
	{
		a.POST("/", authentification.Login)
	}
	u := router.Group("/user")
	{
		u.POST("/get", user.Get_user_by_id_or_name)
	}
}