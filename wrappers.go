package main

import (
    "net/http"
)

func PostOnly(h handler) handler {

    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "POST" {
            h(w, r)
            return
        }
        http.Error(w, "post only", http.StatusMethodNotAllowed)
    }
}