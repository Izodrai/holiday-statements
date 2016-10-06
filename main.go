package main

import (
	"./lib/db"
	"./lib/handler"
	"./lib/tools"
	"github.com/gin-gonic/gin"
)

func main() {
	
	tools.Init_log(true)
	
	tools.Users = make(map[string]tools.User)
	tools.Users_id = make(map[int64]tools.User)
	tools.Connected_users = make(map[int64]tools.User)
	tools.Admins = make(map[string]tools.User)
	tools.Friends = make(map[int64][]int64)
	
	if err := db.Init_db_connect(); err != nil {
		db.Db_connect.Close()
		tools.Fatal_error(err)
	}
	defer db.Db_connect.Close()
	
	if err := db.Init_system(); err != nil {
		tools.Fatal_error(err)
	}
	
	router := gin.Default()

	handler.Handler(router)

	router.Run(":8080")
}