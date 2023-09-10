package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/cbass2404/hello_world_golang/pkg/config"
	"github.com/cbass2404/hello_world_golang/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
