package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Hikari1149/bookings/internal/config"
	"github.com/Hikari1149/bookings/internal/driver"
	"github.com/Hikari1149/bookings/internal/handlers"
	"github.com/Hikari1149/bookings/internal/helpers"
	"github.com/Hikari1149/bookings/internal/models"
	"github.com/Hikari1149/bookings/internal/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errLog *log.Logger

// main in the main application function
func main() {

	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()
	//
	defer close(app.MailChan)
	fmt.Println("Starting mail listener")
	listenForMail()

	// msg := models.MailData{
	// 	To:       "hikari94854@gmail.com",
	// 	From:     "me@here.com",
	// 	Subuject: "Some subject",
	// 	Congtent: "",
	// }

	// app.MailChan <- msg

	fmt.Println(fmt.Sprintf("starting application on port %s", portNumber))
	// _ = http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)

}

func run() (*driver.DB, error) {

	// session ?
	gob.Register(models.Reservation{})
	gob.Register(models.User{})
	gob.Register(models.Room{})
	gob.Register(models.Restriction{})
	gob.Register(models.RommRestriction{})

	//
	mailChan := make(chan models.MailData)
	app.MailChan = mailChan

	// change this to true when in production
	app.InProduction = false

	// set up logger
	infoLog := log.New(os.Stdin, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errLog := log.New(os.Stdin, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errLog

	//
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	// connect to database
	log.Println("Connecting to Database...")

	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=bookings user=a123 password=")
	if err != nil {
		log.Fatal("Cannot connect to database! Dying...", err)
	}
	log.Println("Connected to Database!")
	//
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
		return nil, err
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)
	render.NewRender(&app)
	helpers.NewHelpers(&app)

	return db, nil
}
