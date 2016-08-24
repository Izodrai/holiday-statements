package templates

import (
	"../tools"
	"github.com/abbot/go-http-auth"
	"html/template"
	"net/http"
)

func TemplateMe(w http.ResponseWriter, r *auth.AuthenticatedRequest, page string, info interface{}) {
	tmpl, err := template.ParseFiles(page)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		tools.FatalError(err.Error())
	}

	tmpl.Execute(w, info)

	tools.Info("Requests to -> ", r.Host, r.URL, " -> Status : ", http.StatusOK, " -> For : ", r.Username)
}

func Template500(w http.ResponseWriter, r *auth.AuthenticatedRequest) {
	tmpl, err := template.ParseFiles("lib/templates/500/500.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		tools.FatalError(err.Error())
	}

	tmpl.Execute(w, nil)

	tools.Info("Requests to -> ", r.Host, r.URL, " -> Status : ", http.StatusInternalServerError, " -> For : ", r.Username)
}