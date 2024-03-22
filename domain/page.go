package domain

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
	filename := "page/" + p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func LoadPage(title string) (*Page, error) {
	filename := "page/" + title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}