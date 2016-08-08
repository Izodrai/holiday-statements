package templates

import (
	"net/http"
	"../tools"
	"html/template"
	"github.com/abbot/go-http-auth"
)

func TemplateMe(w http.ResponseWriter, r *auth.AuthenticatedRequest, page string, info interface{}) {
    tmpl, err := template.ParseFiles(page)

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
	tools.FatalError(err.Error())
    }

    tmpl.Execute(w, info)

    tools.Info("Connection to -> ",r.Host,r.URL," -> Status : ",http.StatusOK," -> For : ",r.Username)
}