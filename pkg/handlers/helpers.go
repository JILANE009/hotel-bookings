package helpers

import (
	"fmt"
	"github.com/JILANE009/hotel-bookings/pkg/config"
	"github.com/JILANE009/hotel-bookings/pkg/models"
	"github.com/JILANE009/hotel-bookings/pkg/render"
	"math/rand"
	"net/http"
)

func Hello() {
	fmt.Println("Hello World")
}

func RandomNumber(n int) int {
	value := rand.Intn(n)
	return value
}

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepository(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, req *http.Request) {
	remoteIP := req.RemoteAddr
	m.App.Session.Put(req.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, req, "home.page.gohtml", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, req *http.Request) {

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello World"

	remoteIP := m.App.Session.GetString(req.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, req, "about.page.gohtml", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) Reservation(w http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(w, req, "reservation.page.gohtml", &models.TemplateData{})
}

func (m *Repository) PostReservation(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Posted a search availability"))
}
