package main

import (
	"net/http"

	"github.com/ManojGowda27/bnb-booking/pkg/config"
	"github.com/ManojGowda27/bnb-booking/pkg/handlers"
	"github.com/go-chi/chi"
)

func routes(app *config.AppConfig)http.Handler {
	mux := chi.NewRouter()

	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handler.Repo.Home)
	mux.Get("/about", handler.Repo.About)
	mux.Get("/generals-quarters", handler.Repo.Generals)
	mux.Get("/majors-suite", handler.Repo.Majors)
	mux.Get("/search-availability", handler.Repo.Availability)
	mux.Get("/make-reservation", handler.Repo.Reservation)
	mux.Get("/contact", handler.Repo.Contact)

	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))


	return mux
} 