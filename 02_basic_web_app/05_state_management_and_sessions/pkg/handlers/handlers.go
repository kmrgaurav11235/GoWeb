package handlers

import (
	"net/http"

	"github.com/kmrgaurav11235/goweb/state_management_and_sessions/pkg/config"
	"github.com/kmrgaurav11235/goweb/state_management_and_sessions/pkg/models"
	"github.com/kmrgaurav11235/goweb/state_management_and_sessions/pkg/render"
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
	render.RenderTemplate(rw, "home.page.tmpl", &models.TemplateData{}) // For now, we are passing empty TemplateData
}

// About is the about page handler
func (m *Repository) About(rw http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["Test"] = "Hello again."

	// send the data to the template
	render.RenderTemplate(rw, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
