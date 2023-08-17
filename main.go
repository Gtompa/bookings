package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Gtompa/bookings/pkg/render"

	"github.com/Gtompa/bookings/pkg/handlers"

	"github.com/Gtompa/bookings/pkg/config"

	"github.com/alexedwards/scs/v2"
)

var app config.AppConfig
var session *scs.SessionManager

const PORT_NUMBER = ":4000"

func main() {

	// change this to true when in production

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template c")
	}

	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	if err != nil {
		fmt.Println("An error occurred:", err)
	}

	srv := http.Server{
		Addr:    PORT_NUMBER,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
