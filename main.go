package main

import (
	"./lib/db"
	"./lib/handler"
	"github.com/gin-gonic/gin"
	"github.com/izodrai/utils/logs"
)

func main() {

	logs.Init_log(true)

	if err := db.Init_db_connect(); err != nil {
		db.Db_connect.Close()
		logs.Fatal_error(err)
	}
	defer db.Db_connect.Close()

	if err := db.Init_system(); err != nil {
		logs.Fatal_error(err)
	}

	router := gin.Default()

	handler.Handler(router)

	router.Run(":8080")
}
