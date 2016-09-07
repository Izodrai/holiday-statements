package events

import (
	tmpl "../../templates"
	"../../tools"
	"github.com/abbot/go-http-auth"
	"net/http"
	"../../db"
)

func Add(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
	
	var user tools.User
	
	info := struct {
		Title        string
		Nav          []string
		Events  []tools.Event
		Added bool
		Error bool
		ErrorMsg string
	}{
		Title: "",
		Nav: tools.GenerateNav(r.Username),
		Events: []tools.Event{},
		Added: false,
		Error: false,
		ErrorMsg: "",
	}
	
	if info.Added, err = addEvent(r, &info.Event, spendingTypes); err != nil {
		tools.Error(err)
		tmpl.Template500(w, r)
		return
	}
	

	tmpl.TemplateMe(w, r, "lib/templates/events/list.html", info)
}

func addEvent(r *auth.AuthenticatedRequest) (bool, error) {
	
	var err error
	
	if addEv := r.PostFormValue("addEv"); addEv == "" {
		return false, nil
	} else if addEv != "Ajouter" {
		return false, errors.New("bad entry for addEv -> "+addEv)
	}
	
	var ev tools.Event
	
	if name := r.PostFormValue("name"); name == "" {
		return false, nil /*errors.New("bad entry for amount*/
	} else {
		ev.Reference = name
	}
	
	return true, nil
}
