package auth

import (
	"encoding/base64"
	"net/http"
	"strings"
	"../config"
)

type Handler func(w http.ResponseWriter, r *http.Request)

func BasicAuth(pass Handler, users map[string]config.User) Handler {

    return func(w http.ResponseWriter, r *http.Request) {

        auth := strings.SplitN(r.Header["Authorization"][0], " ", 2)

        if len(auth) != 2 || auth[0] != "Basic" {
            http.Error(w, "bad syntax", http.StatusBadRequest)
            return
        }

        payload, _ := base64.StdEncoding.DecodeString(auth[1])
        pair := strings.SplitN(string(payload), ":", 2)

        if len(pair) != 2 || !Validate(pair[0], pair[1]) {
            http.Error(w, "authorization failed", http.StatusUnauthorized)
            return
        }

        pass(w, r)
    }
}

func Validate(username, password string) bool {
    if username == "username" && password == "password" {
        return true
    }
    return false
}