package main

import (
	"time"
	"./lib/handler"
	"./lib/tools"
	"github.com/gin-gonic/gin"
) 

func main() {
	
	tools.Init_log(true)

	router := gin.Default()

	tools.Users = make(map[string]tools.User)
	tools.Users_id = make(map[int64]tools.User)
	tools.Connected_users = make(map[int64]tools.User)
	
	tools.Users["test"]=tools.User{1,"test","a1159e9df3670d549d04524532629f5477ceb7deec9b45e47e8c009506ecb2c8","@",false,[]int64{},"",time.Time{}}
	tools.Users_id[1]=tools.User{1,"test","a1159e9df3670d549d04524532629f5477ceb7deec9b45e47e8c009506ecb2c8","@",false,[]int64{},"",time.Time{}}
	
	
	
	
	
	
	
	
	
	
	
	handler.Handler(router)
	
	router.Run(":8080")
}