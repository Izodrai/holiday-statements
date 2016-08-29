package events

import (
	tmpl "../../templates"
	"../../tools"
	"github.com/abbot/go-http-auth"
	"net/http"
	"strconv"
	"../../db"
	"time"
)

func Get(w http.ResponseWriter, r *auth.AuthenticatedRequest) {

	var err error
	info := struct {
		Title        string
		Nav          []string
		Actualize    int64
		Event tools.Event
		Date string
		TypeSpending []tools.SpendingType
	}{
		Title: "évènement",
		Nav: tools.GenerateNav(r.Username),
		Actualize: 0,
                Event: tools.Event{},
		Date: time.Now().Format("2006-01-02"),
		TypeSpending: []tools.SpendingType{},
	}
	
	params := r.URL.Query()
	
	p, _ := params["get"]
	
	info.Event.Id, err = strconv.ParseInt(p[0],10,64)
	if err != nil {
		tools.Error(err)
		tmpl.Template500(w, r)
		return
	}
	
	user, _ := tools.Users[r.Username]
	
	ok, err := db.CheckEventForThisUser(&user, &info.Event)
	if err != nil {
		tools.Error(err)
		tmpl.Template500(w, r)
		return
	}
	
	if !ok {
		http.Redirect(w, &r.Request, "/events", http.StatusForbidden)
		return
	}
	
	//TODO add events if exist
	
	if err = db.LoadThisEvent(&info.Event); err != nil {
		tools.Error(err)
		tmpl.Template500(w, r)
		return
	}
	
	if err = db.LoadTypeSpending(&info.TypeSpending); err != nil {
		tools.Error(err)
		tmpl.Template500(w, r)
		return
	}
	
	info.Title = info.Event.Reference
	info.Actualize = info.Event.Id
	
	tmpl.TemplateMe(w, r, "lib/templates/events/get.html", info)
}