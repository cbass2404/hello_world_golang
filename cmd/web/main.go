package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/cbass2404/hello_world_golang/pkg/config"
	"github.com/cbass2404/hello_world_golang/pkg/handlers"
	"github.com/cbass2404/hello_world_golang/pkg/render"
)

const portNumber = ":8080"

var app config.AppConfig

// main is the main application function
func main() {

	// change this to true when in production
	app.InProduction = false

	app.Session = scs.New()
	app.Session.Lifetime = 24 * time.Hour
	app.Session.Cookie.Persist = true
	app.Session.Cookie.SameSite = http.SameSiteLaxMode
	app.Session.Cookie.Secure = app.InProduction

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
