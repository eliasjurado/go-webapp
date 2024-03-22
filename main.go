package main

import (
	"gowiki/handler"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/view/", handler.ViewHandler)
	http.HandleFunc("/edit/", handler.EditHandler)
	http.HandleFunc("/save/", handler.SaveHandler)
	
	log.Fatal(http.ListenAndServe(":8080", nil))
}
