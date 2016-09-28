package authentification

import (
	"github.com/gin-gonic/gin"
)


func Handler(router *gin.Engine) {
	authentification := router.Group("/authentification")
	{
		authentification.POST("/", login)
	}
}