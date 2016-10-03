package main

import (
	"./lib/db"
	"./lib/handler"
	"./lib/tools"
	"./lib/authentification"
	"github.com/gin-gonic/gin"
) 

func main() {
	
	tools.Init_log(true)
	
	if err := db.Init_connect_and_db(); err != nil {
		db.Db_connect.Close()
		tools.Fatal_error(err)
	}

	if err := authentification.Init_authenfication(); err != nil {
		tools.Fatal_error(err)
	}
	
	router := gin.Default()

	handler.Handler(router)
	
	router.Run(":8080")
}