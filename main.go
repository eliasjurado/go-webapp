package main

import (
	"fmt"
	"gowiki/page"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := page.LoadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1> <div>%s</div>", p.Title, p.Body)
}
