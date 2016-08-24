package events

import (
	tmpl "../../templates"
	"../../tools"
	"github.com/abbot/go-http-auth"
	"net/http"
)

func Get(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
	
	nav := tools.GenerateNav(r.Username)
	
	info := struct {
		Title        string
		Nav          []string
		Participants []string
	}{
		Title: "get events",
		Nav: nav,
		Participants: []string{
			"Valentin",
			"Emma",
			"Justine",
		},
	}

	tmpl.TemplateMe(w, r, "lib/templates/events/get.html", info)
}