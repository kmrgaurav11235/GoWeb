package handlers

import (
	"net/http"

	"github.com/kmrgaurav11235/goweb/more_on_templates/pkg/render"
)

// Home is the home page handler
func Home(rw http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(rw, "home.page.tmpl")
}

// About is the about page handler
func About(rw http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(rw, "about.page.tmpl")
}
