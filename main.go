package main

import (
	localAuth "./lib/auth"
	"./lib/db"
	"./lib/routes/events"
	"./lib/routes/index"
	"./lib/routes/users"
	"./lib/tools"
	auth "github.com/abbot/go-http-auth"
	"log"
	"path"
	"net/http"
	"os"
)

func main() {

	var err error
	var config tools.Config

	if len(os.Args) != 2 {
		log.Println(tools.RED+ "Invalid Argument(s)"+ tools.STOP)
		log.Println(tools.RED+ "Usuel : ./holi-binary config_file"+ tools.STOP)
		os.Exit(1)
	}

	if config, err = tools.LoadConfig(path.Join(os.Args[1])); err != nil {
		log.Println(tools.RED+err.Error()+tools.STOP)
		os.Exit(1)
	}

	tools.InitLog(true)

	if err = db.Init(config); err != nil {
		db.DbConnect.Close()
		tools.FatalError(err)
	}
	defer db.DbConnect.Close()

	if err := localAuth.Init(); err != nil {
		tools.FatalError(err)
	}

	userAuth := auth.NewBasicAuthenticator("Current Authentication", localAuth.UserSecret)
	adminAuth := auth.NewBasicAuthenticator("Admin Authentication", localAuth.AdminSecret)

	http.HandleFunc("/", index.HandleDefault)

// 	http.HandleFunc("/index", userAuth.Wrap(index.HandleIndex))

	http.HandleFunc("/events", userAuth.Wrap(events.HandleEvents))

	http.HandleFunc("/users", adminAuth.Wrap(users.HandleUsers))

	tools.Info("Working")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
