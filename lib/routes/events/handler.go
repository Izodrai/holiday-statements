package events

import (
	"github.com/abbot/go-http-auth"
	"net/http"
	
	evs "../../workers/events"
)

func HandleEvents(w http.ResponseWriter, r *auth.AuthenticatedRequest) {

	params := r.URL.Query()
	
	if _, ok := params["get"]; ok {
		evs.Get(w,r)
		return 
	}
	
	evs.List(w,r)
}
