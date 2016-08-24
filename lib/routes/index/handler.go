package index

import (
	"net/http"
)

func HandleDefault(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/events", 301)
}