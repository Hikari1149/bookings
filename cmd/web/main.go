package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Hikari1149/go-web-example/pkg/config"
	"github.com/Hikari1149/go-web-example/pkg/handlers"
	"github.com/Hikari1149/go-web-example/pkg/render"
)

const portNumber = ":8080"

// main in the main application function
func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("starting application on port %s", portNumber))
	// _ = http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)

}
