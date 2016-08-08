package index

import (
	tmpl "../../templates"
	"../../tools"
	"github.com/abbot/go-http-auth"
	"net/http"
)

func HandleDefault(w http.ResponseWriter, r *http.Request) {
	tools.Info("/")
	http.Redirect(w, r, "/index", 301)
}

func HandleIndex(w http.ResponseWriter, r *auth.AuthenticatedRequest) {

	tools.Info(r.URL.Query())
	
// 	r.ParseForm()
// 	payeur, _ = r.PostForm["payeur"]
	
// 	test1, _ := r.PostForm["test1"]
// 	test2, _ := r.PostForm["test2"]
// 	
// 	tools.Info("test1 : ", test1)
// 	tools.Info("test2 : ", test2)
	
	info := struct {
		Title        string
		Participants []string
	}{
		Title: "index",
		Participants: []string{
			"Valentin",
			"Emma",
			"Justine",
			"Jérôme",
			"Vincent",
		},
	}

	tmpl.TemplateMe(w, r, "lib/templates/index/index.html", info)
}
