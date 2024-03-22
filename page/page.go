package page

import "os"

type Page struct {
	Title string
	Body  []byte
}

func NewPage(
	title string,
	body string,
) (obj *Page, err error) {
	obj = &Page{
		Title: title,
		Body:  []byte(body),
	}
	return
}

func (p *Page) Save() error {
	filename := "pages/" + p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}