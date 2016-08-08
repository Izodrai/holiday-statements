package index

import (
	"net/http"
	"../../tools"
	tmpl "../../templates"
	"github.com/abbot/go-http-auth"
)

func HandleDefault(w http.ResponseWriter, r *http.Request) {
	tools.Info("/")
	http.Redirect(w, r, "/index", 301)
}

func HandleIndex(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
	
	info := struct {
		Title string
		Participants []string
	}{
		Title : "index",
		Participants : []string{
			"Valentin",
			"Emma",
			"Justine",
			"Jérôme",
			"Vincent",
		},
	}
	
	tmpl.TemplateMe(w, r, "lib/templates/index/index.html", info)
}