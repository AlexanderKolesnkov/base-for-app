package main

import (
	"github.com/AlexanderKolesnkov/WebHotel/pkg/config"
	"github.com/AlexanderKolesnkov/WebHotel/pkg/handlers"
	"github.com/bmizerany/pat"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}