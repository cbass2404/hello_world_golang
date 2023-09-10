package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/cbass2404/hello_world_golang/pkg/config"
	"github.com/cbass2404/hello_world_golang/pkg/handlers"
	"github.com/cbass2404/hello_world_golang/pkg/render"
)

const portNumber = ":8080"

// main is the main application function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.UseCache = false
	app.TemplateCache = tc

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	fmt.Printf("Starting application on port %s", portNumber)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
