package page

import (
	"os"
	"io"
	"time"
	"bytes"
	"text/template"
	"log"
)

type Page struct {
	pageName string
	templateName string

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

func NewPage(templateName string) (p *Page) {
	p = new(Page)
	p.Now = time.Now()
	p.templateName = templateName
	return
}

func (p *Page) Execute(w io.Writer) {
	var err error
	p.template, err = template.ParseFiles("template/" + p.templateName)
	if err == nil {
		p.template.Execute(w, p)
	}else{
		log.Fatal(err)
	}
}

func (p *Page) FromFile(slug string) (*Page) {
	body, _ := template.ParseFiles("page/" + slug)
	var b bytes.Buffer
	body.Execute(&b, p)
	p.Body = b.String()
	return p
}

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

