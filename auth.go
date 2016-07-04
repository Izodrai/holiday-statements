package main

import (
	"encoding/base64"
	"net/http"
	"strings"
	"encoding/json"
	"io/ioutil"
	"./lib/tools"
)

type handler func(w http.ResponseWriter, r *http.Request)

func BasicAuth(pass handler, users map[string]User) handler {

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


type User struct {
	user   string `json:"user"`
	pwd    string `json:"pwd"`
	rights int    `json:"rights"`
}

func LoadUsers(users map[string]User) error {
	
	tools.Info("1")
	
	usersConfigFile, err := ioutil.ReadFile("config/users.json")
	if err != nil {
		return err
	}
	tools.Info("2")

	if err := json.Unmarshal([]byte(usersConfigFile), &users); err != nil {
		return err
	}
	tools.Info("3")
	
	return nil
}