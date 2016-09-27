package main

import (
	"./lib/login"
	"./lib/users"
	"./lib/tools"
	"github.com/gin-gonic/gin"
) 

func main() {
	
	tools.InitLog(true)

	router := gin.Default()

	login.Handler(router)
	users.Handler(router)
	
	router.Run(":8080")
}