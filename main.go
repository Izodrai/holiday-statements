package main

import (
	"time"
	"./lib/tools"
	"github.com/gin-gonic/gin"
)

func main() {

	tools.InitLog(true)
	
	tools.Info("Working")
	
	router := gin.New()
	
	router.LoadHTMLFiles("lib/templates/index.templ.html", "lib/templates/login.templ.html")
	
	router.GET("/", index)
	router.GET("/login", login)
	
	router.Run(":8080")
}

func index(c *gin.Context) {
	c.Redirect(301, "/login")
}

func login(c *gin.Context) {
	c.HTML(200, "login.templ.html", gin.H{
		"timestamp": time.Now().Unix(),
	})
}