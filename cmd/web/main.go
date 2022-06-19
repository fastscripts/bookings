package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/fastscripts/bookings/pkg/config"
	"github.com/fastscripts/bookings/pkg/handlers"
	"github.com/fastscripts/bookings/pkg/render"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main programm
func main() {

	//change this to true when in Production
	app.InProduction = false

	session = scs.New()

	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = templateCache
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	/* 	http.HandleFunc("/", handlers.Repo.Home)
	   	http.HandleFunc("/about", handlers.Repo.About) */

	fmt.Printf("Starting application on port %s \n", portNumber)
	// _ = http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
