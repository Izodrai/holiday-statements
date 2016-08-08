package events

import (
	"../../tools"
	"fmt"
	"github.com/abbot/go-http-auth"
	"net/http"
)

func HandleEvents(w http.ResponseWriter, r *auth.AuthenticatedRequest) {

	tools.Info("events")

	fmt.Fprint(w, `<html>`+tools.Menu+`<p>events</p>`, r.Username)
}
