package main

import (
// 	"log"
// 	"net/http"
	"./lib/tools"
)

func main() {

	tools.InitLog(true)
	
	var err error
	var users = make(map[string]User)
	
	if err = LoadUsers(users); err != nil {
		tools.FatalError(err)
		return
	}
	
	tools.Info(users)
	
// 	// public views
// 	http.HandleFunc("/", HandleIndex)
// 
// 	// private views
// 	http.HandleFunc("/post", PostOnly(BasicAuth(HandlePost,users)))
// 
// 	log.Fatal(http.ListenAndServe(":8080", nil))
}