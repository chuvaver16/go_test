package handler

import (
	"errors"
	"io"
	"log"
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
)

type TemplateRegistry struct {
	Templates map[string]*template.Template
}

// Implement e.Renderer interface
func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.Templates[name]
	if !ok {
		err := errors.New("Template not found -> " + name)
		return err
	}
	return tmpl.ExecuteTemplate(w, "index.html", data)
}

func InitTemplates() *TemplateRegistry {
	t := new(TemplateRegistry)
	t.Templates = make(map[string]*template.Template)
	t.Templates["weather.html"] = template.Must(template.ParseFiles("view/weather.html", "view/index.html"))
	return t
}

func MainHandler(c echo.Context) error {
	log.Println("Here")

	return c.Render(http.StatusOK, "weather.html", nil)
}
