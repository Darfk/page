package page

import (
	"bytes"
	"io"
	_ "log"
	"path/filepath"
	"text/template"
	"time"
)

type Page struct {
	body     *template.Template
	template *template.Template

	Head        []string
	Scripts     []string
	Stylesheets []string
	Title       string
	Description string
	Keywords    []string
	Body        string
	Now         time.Time
}

func NewPage() (p *Page) {
	p = new(Page)
	p.Now = time.Now()
	return
}

func (p *Page) LoadTemplate(templateFile string) (err error) {
	p.template, err = template.ParseFiles(filepath.Join("template", templateFile))
	return
}

func (p *Page) LoadBody(bodyFile string) (err error) {
	p.body, err = template.ParseFiles(filepath.Join("page", bodyFile))
	return
}

func (p *Page) Execute(w io.Writer) (err error) {

	if p.template != nil && p.body == nil {
		err = p.template.Execute(w, p)
		return
	}

	if p.template == nil && p.body != nil {
		err = p.body.Execute(w, p)
		return
	}

	if p.template != nil && p.body != nil {
		var b bytes.Buffer
		err = p.body.Execute(&b, p)
		if err != nil {
			return
		}
		p.Body = b.String()
		err = p.template.Execute(w, p)
		return
	}

	return
}

func (p *Page) SetTitle(title string) (s string) {
	p.Title = title
	return
}

func (p *Page) AddJS(path string) (s string) {
	p.Scripts = append(p.Scripts, path)
	return
}

func (p *Page) UseTemplate(path string) (s string) {
	p.LoadTemplate(path)
	return
}

/*

func Exists(slug string) bool {
	fh, err := os.Open("page/" + slug)
	defer fh.Close()
	if err == nil {
		return true
	}
	return false
}

func TemplateExists(slug string) bool {
	fh, err := os.Open("template/" + slug)
	defer fh.Close()
	if err == nil {
		return true
	}
	return false
}

*/
