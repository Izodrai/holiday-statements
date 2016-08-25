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

	var err error
	info := struct {
		Title        string
		Nav          []string
	}{
		Title: "évènement",
		Nav: tools.GenerateNav(r.Username),
	}
	
	params := r.URL.Query()
	
	p, _ := params["get"]
	
	var ev tools.Event
	
	ev.Id, err = strconv.ParseInt(p[0],10,64)
	if err != nil {
		tools.Error(err)
		tmpl.Template500(w, r)
		return
	}
	
	user, _ := tools.Users[r.Username]
	
	ok, err := db.CheckEventForThisUser(&user, &ev)
	if err != nil {
		tools.Error(err)
		tmpl.Template500(w, r)
		return
	}
	
	if !ok {
		http.Redirect(w, &r.Request, "/events", http.StatusForbidden)
		return
	}
	
	if err = db.LoadThisEvent(&ev); err != nil {
		tools.Error(err)
		tmpl.Template500(w, r)
		return
	}
	
	info.Title = ev.Reference
	

	tmpl.TemplateMe(w, r, "lib/templates/events/get.html", info)
}