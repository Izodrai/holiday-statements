package main

import (
	"time"
	"./lib/login"
	"./lib/users"
	"./lib/tools"
	"github.com/gin-gonic/gin"
) 

func main() {
	
	tools.Init_log(true)

	router := gin.Default()

	users.Users = make(map[string]users.User)
	users.Users_id = make(map[int64]users.User)
	users.Connected_users = make(map[int64]users.User)
	
	users.Users["test"]=users.User{1,"test","a1159e9df3670d549d04524532629f5477ceb7deec9b45e47e8c009506ecb2c8","@",false,[]int64{},"",time.Time{}}
	users.Users_id[1]=users.User{1,"test","a1159e9df3670d549d04524532629f5477ceb7deec9b45e47e8c009506ecb2c8","@",false,[]int64{},"",time.Time{}}
	
	
	
	
	login.Handler(router)
	users.Handler(router)
	
	router.Run(":8080")
}