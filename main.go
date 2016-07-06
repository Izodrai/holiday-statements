package main

import (
	"os"
	"log"
	"net/http"
	"./lib/tools"
	"./lib/routes"
)

func main() {

	tools.InitLog(true)
	
	tools.Info("Working")
	
	err := os.Remove("save.csv")
	if err != nil {
		if !os.IsNotExist(err) {
			panic(err)
		}
	}
	
	f, err := os.OpenFile("save.csv", os.O_APPEND|os.O_CREATE|os.O_RDWR, os.ModeAppend|0755)
	if err != nil {
		panic(err)
	}
	
	if _, err := f.Write([]byte("date;montant;by;emma;;justine;;valentin;;jerome;;Nfor\n")); err != nil {
		panic(err)
	}
	
	f.Close()
	
	http.HandleFunc("/", routes.HandleIndex)

	log.Fatal(http.ListenAndServe(":8080", nil))
}