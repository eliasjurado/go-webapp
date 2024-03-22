package main

import (
	"gowiki/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/view/", handler.MakeHandler(handler.ViewHandler))
    http.HandleFunc("/edit/", handler.MakeHandler(handler.EditHandler))
    http.HandleFunc("/save/", handler.MakeHandler(handler.SaveHandler))
	
	log.Fatal(http.ListenAndServe(":8080", nil))
}
