package page

import (
	"bytes"
	"github.com/russross/blackfriday"
	"io"
	"path/filepath"
	"text/template"
	"time"
)

var TemplateRoot string = "template"
var PageRoot string = "page"

type Meta struct {
	Head        []string
	Scripts     []string
	Stylesheets []string
	Title       string
	Description string
	Keywords    []string
	Body        string
	page        *Page
	format      string
}

type Page struct {
	body     *template.Template
	template *template.Template
	meta     *Meta
}

func NewPage() (p *Page) {
	p = new(Page)
	p.meta = new(Meta)
	p.meta.page = p
	p.meta.format = "markdown"
	return
}

func (p *Page) Execute(w io.Writer) (err error) {

	if p.template != nil && p.body == nil {
		err = p.template.Execute(w, p)
		return
	}

	if p.template == nil && p.body != nil {
		var b bytes.Buffer
		err = p.body.Execute(&b, p)
		if err != nil {
			return
		}

		w.Write(blackfriday.MarkdownCommon(b.Bytes()))
		return
	}

	if p.template != nil && p.body != nil {

		var b bytes.Buffer
		err = p.body.Execute(&b, p.meta)
		if err != nil {
			return
		}

		if p.meta.format == "markdown" {
			p.meta.Body = string(blackfriday.MarkdownCommon(b.Bytes()))
		} else {
			p.meta.Body = string(b.Bytes())
		}

		err = p.template.Execute(w, p.meta)

		return
	}

	return
}

func (p *Page) LoadTemplate(templateFile string) (err error) {
	p.template, err = template.ParseFiles(filepath.Join(TemplateRoot, templateFile))
	return
}

func (p *Page) LoadBody(bodyFile string) (err error) {
	p.body, err = template.ParseFiles(filepath.Join(PageRoot, bodyFile))
	return
}

func (m *Meta) SetFormat(format string) (s string) {
	m.format = format
	return
}

func (m *Meta) Now() time.Time {
	return time.Now()
}

func (m *Meta) SetTitle(title string) (s string) {
	m.Title = title
	return
}

func (m *Meta) AddJS(path string) (s string) {
	m.Scripts = append(m.Scripts, path)
	return
}

func (m *Meta) AddCSS(path string) (s string) {
	m.Stylesheets = append(m.Stylesheets, path)
	return
}
