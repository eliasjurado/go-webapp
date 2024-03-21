package main

import (
	"gowiki/page"
	"log"
)

func main() {
	log.SetFlags(0)
	p1, _ := page.NewPage("TestPage", "Esta es una página de muestra.")
	p1.Save()

	p2, _ := page.LoadPage("TestPage")
	log.Printf("%+v\n", string(p2.Body))

}
