package index

import (
// 	tmpl "../../templates"
// 	"../../tools"
// 	"github.com/abbot/go-http-auth"
	"net/http"
)

func HandleDefault(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/events", 301)
}

// func HandleIndex(w http.ResponseWriter, r *auth.AuthenticatedRequest) {

// 	tools.Info(r.URL.Query())
	
// 	r.ParseForm()
// 	payeur, _ = r.PostForm["payeur"]
	
// 	test1, _ := r.PostForm["test1"]
// 	test2, _ := r.PostForm["test2"]
// 	
// 	tools.Info("test1 : ", test1)
// 	tools.Info("test2 : ", test2)
	
	
// 	nav := tools.GenerateNav(r.Username)
// 	
// 	info := struct {
// 		Title        string
// 		Nav          []string
// 		Participants []string
// 	}{
// 		Title: "index",
// 		Nav: nav,
// 		Participants: []string{
// 			"Valentin",
// 			"Emma",
// 			"Justine",
// 		},
// 	}
// 
// 	tmpl.TemplateMe(w, r, "lib/templates/index/index.html", info)
// }
