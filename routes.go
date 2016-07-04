
package main

import (
    "io"
    "log"
    "net/http"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "hello, world\n")
}

func HandlePost(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()
    log.Println(r.PostForm)
    io.WriteString(w, "post\n")
}

type Result struct {
    FirstName string `json:"first"`
    LastName  string `json:"last"`
}
