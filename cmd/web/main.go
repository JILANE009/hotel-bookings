package main

import (
	"fmt"
	"github.com/JILANE009/hotel-bookings/internal/config"
	"github.com/JILANE009/hotel-bookings/internal/handlers"
	"github.com/JILANE009/hotel-bookings/internal/render"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"
)

const portName = "0.0.0.0:4504"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	// change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := helpers.NewRepository(&app)
	helpers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println("Listening on port " + portName)

	srv := http.Server{
		Addr:    portName,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
