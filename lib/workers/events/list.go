package events

import (
	tmpl "../../templates"
	"../../tools"
	"github.com/abbot/go-http-auth"
	"net/http"
	"../../db"
)

func List(w http.ResponseWriter, r *auth.AuthenticatedRequest) {

	var user tools.User

	info := struct {
		Title        string
		Nav          []string
		Events  []tools.Event
	}{
		Title: "Mes évènements",
		Nav: tools.GenerateNav(r.Username),
		Events: []tools.Event{},
	}

	user, _ = tools.Users[r.Username]
	if err := db.LoadEventsForThisUser(&user, &info.Events); err != nil {
		tools.Error(err)
		tmpl.Template500(w, r)
		return
	}

	tmpl.TemplateMe(w, r, "events/list.html", info)
}
