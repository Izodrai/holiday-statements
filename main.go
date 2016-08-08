package main

import (
// 	"log"
// 	"net/http"
	"./lib/db"
	"./lib/tools"
	localAuth "./lib/auth"
// 	"./lib/routes/index"
// 	"./lib/routes/events"
// 	"./lib/routes/users"
// 	auth "github.com/abbot/go-http-auth"
)



func main() {
	
// 	userAuth := auth.NewBasicAuthenticator("Current Authentication", localAuth.UserSecret)
// 	adminAuth := auth.NewBasicAuthenticator("Admin Authentication", localAuth.AdminSecret)
// 	
// 	http.HandleFunc("/", index.HandleDefault)
// 	
// 	http.HandleFunc("/index", userAuth.Wrap(index.HandleIndex))
// 	
// 	http.HandleFunc("/events", userAuth.Wrap(events.HandleEvents))
// 	
// 	http.HandleFunc("/users", adminAuth.Wrap(users.HandleUsers))

	tools.Info("Working")
	
// 	log.Fatal(http.ListenAndServe(":8080", nil))
}



func init() {
	tools.InitLog(true)
	
	if err := db.Init(); err != nil {
		tools.FatalError(err)
	}
	
	if err := localAuth.Init(); err != nil {
		tools.FatalError(err)
	}
	
	
	
}