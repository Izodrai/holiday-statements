package login

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

// http://localhost:8080/login/

// curl -i -X POST -d '{"login":"C-3PO","pwd":"R2D2"}' http://localhost:8080/login/

func authentificate(c *gin.Context) {

	var json login_form
	
	if c.BindJSON(&json) == nil {
		fmt.Println(json.Login)
		fmt.Println(json.Pwd)
		
		c.JSON(http.StatusOK, gin.H{
			"login": json.Login,
			"pwd": json.Pwd,
			"status": "you are logged in",
		})
	}
}

type login_form struct {
    Login     string `form:"login" binding:"required"`
    Pwd string `form:"pwd" binding:"required"`
}