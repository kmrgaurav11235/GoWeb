package handlers

import (
	"net/http"

	"github.com/kmrgaurav11235/goweb/working_with_forms/pkg/config"
	"github.com/kmrgaurav11235/goweb/working_with_forms/pkg/models"
	"github.com/kmrgaurav11235/goweb/working_with_forms/pkg/render"
)

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// Repo the Repository used by the handlers
var Repo *Repository

// NewRepo creates a new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(rw http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr                                     // Pulls the IP of the visitor
	m.App.SessionManager.Put(r.Context(), "remote_ip", remoteIp) // Put the IP in the SessionManager

	render.RenderTemplate(rw, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(rw http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["Test"] = "Hello again."

	remoteIp := m.App.SessionManager.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp

	// send the data to the template
	render.RenderTemplate(rw, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
