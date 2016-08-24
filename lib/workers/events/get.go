package events

import (
	tmpl "../../templates"
	"../../tools"
	"github.com/abbot/go-http-auth"
	"net/http"
	"strconv"
	"../../db"
)

func Get(w http.ResponseWriter, r *auth.AuthenticatedRequest) {

	info := struct {
		Title        string
		Nav          []string
	}{
		Title: "évènement",
		Nav: tools.GenerateNav(r.Username),
	}
	
	params := r.URL.Query()
	
	p, _ := params["get"]
	
	evId, err := strconv.Atoi(p[0])
	if err != nil {
		tools.Error(err)
		tmpl.Template500(w, r)
		return
	}
	tools.Info(evId)
	
	user, _ := tools.Users[r.Username]
	
	ok, err := db.CheckEventForThisUser(&user, evId)
	if err != nil {
		tools.Error(err)
		tmpl.Template500(w, r)
		return
	}
	
	if !ok {
		http.Redirect(w, &r.Request, "/events", http.StatusForbidden)
		return
	}
	

	tmpl.TemplateMe(w, r, "lib/templates/events/get.html", info)
}