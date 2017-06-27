package events

import (
	tmpl "../../templates"
	"../../tools"
	"github.com/abbot/go-http-auth"
	"net/http"
	"../../db"
	"errors"
	"strconv"
)

func Add(w http.ResponseWriter, r *auth.AuthenticatedRequest) {

	var err error

	info := struct {
		Title        string
		Nav          []string
		Event  tools.Event
		Users  []tools.User
		Added bool
		Error bool
		ErrorMsg string
	}{
		Title: "",
		Nav: tools.GenerateNav(r.Username),
		Event: tools.Event{},
		Users:  []tools.User{},
		Added: false,
		Error: false,
		ErrorMsg: "",
	}

	var users = make(map[string]tools.User)

	if err = db.LoadUsers(users); err != nil {
		tools.Error(err)
		tmpl.Template500(w, r)
		return
	}

	for _,user := range users {
		info.Users = append(info.Users, user)
	}

	if info.Added, err = addEvent(r, &info.Event, info.Users); err != nil {
		tools.Error(err)
		tmpl.Template500(w, r)
		return
	}

	tmpl.TemplateMe(w, r, "events/add.html", info)
}

func addEvent(r *auth.AuthenticatedRequest, ev *tools.Event, users []tools.User) (bool, error) {

	var err error

	if addEv := r.PostFormValue("addEv"); addEv == "" {
		return false, nil
	} else if addEv != "Ajouter" {
		return false, errors.New("bad entry for addEv -> "+addEv)
	}

	if name := r.PostFormValue("name"); name == "" {
		return false, nil /*errors.New("bad entry for amount*/
	} else {
		ev.Reference = name
	}

	if allParticipant := r.PostFormValue("allParticipant"); allParticipant == "on" {
		for _, user := range users {
			ev.Participants = append(ev.Participants, user)
		}
	} else {
		for _, user := range users {
			if participe := r.PostFormValue(strconv.FormatInt(user.Id, 10)+"-Participant"); participe == "on" {
				ev.Participants = append(ev.Participants, user)
			}
		}
	}
	promoter := tools.Users[r.Username]
	ev.PromoterId = promoter.Id

	if len(ev.Participants) == 0 {
		return false, nil
	}

	if err = db.AddThisEvent(ev); err != nil {
		return false, err
	}

	return true, nil
}
