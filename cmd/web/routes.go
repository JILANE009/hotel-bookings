package main

import (
	"github.com/JILANE009/hotel-bookings/pkg/config"
	helpers "github.com/JILANE009/hotel-bookings/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", helpers.Repo.Home)
	mux.Get("/about", helpers.Repo.About)
	mux.Get("/reservation", helpers.Repo.Reservation)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
