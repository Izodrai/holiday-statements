package main

import (
	"./lib/users"
	"./lib/tools"
	"github.com/gin-gonic/gin"
) 

func main() {
	
	tools.InitLog(true)

	router := gin.Default()

	users.Handler(router)
	
	router.Run(":8080")
}