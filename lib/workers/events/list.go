package events

import (
	tmpl "../../templates"
	"../../tools"
	"github.com/abbot/go-http-auth"
	"net/http"
	"../../db"
)

func List(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
	
	var evs []tools.Event
	var user tools.User
	nav := tools.GenerateNav(r.Username)
	user, _ = tools.Users[r.Username]
	
	info := struct {
		Title        string
		Nav          []string
		Participants []string
	}{
		Title: "list events",
		Nav: nav,
		Participants: []string{
			"Valentin",
			"Emma",
			"Justine",
		},
	}
	
	if err := db.LoadEventsForThisUser(&user, &evs); err != nil {
		tools.Error(err)
		tmpl.TemplateMe(w, r, "lib/templates/events/list.html", info)
		return
	}
	
	tools.Info(evs)
	
	

	tmpl.TemplateMe(w, r, "lib/templates/events/list.html", info)
}

