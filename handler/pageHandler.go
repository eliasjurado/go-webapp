package handler

import (
	"gowiki/domain"
	"html/template"
	"net/http"
)

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := domain.LoadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
	}
	tp := "view"
	RenderView(tp, w, p)
}

func EditHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := domain.LoadPage(title)
	if err != nil {
		p = &domain.Page{Title: title}
	}
	tp := "edit"
	RenderView(tp, w, p)
}

func SaveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &domain.Page{Title: title, Body: []byte(body)}
	p.Save()
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func RenderView(tmplate string, w http.ResponseWriter, p *domain.Page) {
	tp := "template/" + tmplate + ".html"
	t, err := template.ParseFiles(tp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, p)
}
