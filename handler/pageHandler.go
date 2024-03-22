package handler

import (
	"gowiki/domain"
	"html/template"
	"net/http"
	"regexp"
)

func ViewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := domain.LoadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
	}
	tp := "view"
	RenderView(tp, w, p)
}

func EditHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := domain.LoadPage(title)
	if err != nil {
		p = &domain.Page{Title: title}
	}
	tp := "edit"
	RenderView(tp, w, p)
}

func SaveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &domain.Page{Title: title, Body: []byte(body)}
	err := p.Save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func RenderView(tmplate string, w http.ResponseWriter, p *domain.Page) {
	tp := "template/" + tmplate + ".html"
	t, err := template.ParseFiles(tp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

//funcion closure que hace un wrap de otras funciones y valores fuera de ella... hablamos de facade?
func MakeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        m := validPath.FindStringSubmatch(r.URL.Path)
        if m == nil {
            http.NotFound(w, r)
            return
        }
        fn(w, r, m[2])
    }
}