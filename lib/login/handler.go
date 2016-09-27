package login

import (
	"github.com/gin-gonic/gin"
)


func Handler(router *gin.Engine) {
	
	login := router.Group("/login")
	{
		login.POST("/", authentificate)
	}
}