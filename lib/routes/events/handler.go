package events

import (
	"fmt"
	"net/http"
	"../../tools"
	"github.com/abbot/go-http-auth"
)

func HandleEvents(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
	
	tools.Info("events")
	
	fmt.Fprint(w,`<html>`+tools.Menu+`<p>events</p>`,r.Username)
}