package render

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{} // This will allow us to create our own functions and pass them to the template

// RenderTemplate renders templates using html/template
func RenderTemplate(rw http.ResponseWriter, tmpl string) {
	_, err := RenderTemplateTest(rw)
	if err != nil {
		fmt.Println("Error getting template cache:", err)
	}

	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)

	err = parsedTemplate.Execute(rw, nil)

	if err != nil {
		fmt.Println("error parsing template:", err)
	}
}

func RenderTemplateTest(rw http.ResponseWriter) (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{} // map of string -> *Template

	// Get file-names in "templates" folder matching "*.page.tmpl"
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page) // Extract the filename from the entire path

		fmt.Println("Page is currently", page)

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
