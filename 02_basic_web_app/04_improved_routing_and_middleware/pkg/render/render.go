package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/kmrgaurav11235/goweb/improved_routing_and_middleware/pkg/config"
	"github.com/kmrgaurav11235/goweb/improved_routing_and_middleware/pkg/models"
)

var functions = template.FuncMap{} // This will allow us to create our own functions and pass them to the template

var app *config.AppConfig

// NewTemplates sets the config for the render package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDedafultData(td *models.TemplateData) *models.TemplateData {
	// We can add some default template data here
	return td
}

// RenderTemplate renders templates using html/template
func RenderTemplate(rw http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	td = AddDedafultData(td)

	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(rw)

	if err != nil {
		fmt.Println("error writing template to browser:", err)
	}
}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{} // map of string -> *Template

	// Get file-names in "templates" folder matching "*.page.tmpl"
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page) // Extract the filename from the entire path

		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// Get file-names in "templates" folder matching "*.layout.tmpl"
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
