package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/CloudyKit/jet/v6"
	"github.com/alexedwards/scs/v2"
)

type application struct {
	appName string
	server  server
	debug   bool
	errLog  *log.Logger
	infoLog *log.Logger
	view    *jet.Set
	session *scs.SessionManager
}

type server struct {
	host string
	port string
	url  string
}

func main() {

	server := server{
		host: "localhost",
		port: "8080",
		url:  "http://localhost:8080",
	}
	app := &application{
		server:  server,
		appName: "Hnews",
		debug:   true,
		infoLog: log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		errLog:  log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Llongfile),
	}

	// with jet, you can render templates with data
	// with jet.InDevelopmentMode is true, auto reload templates
	if app.debug {
		app.view = jet.NewSet(jet.NewOSFileSystemLoader("./views"), jet.InDevelopmentMode())
	} else {
		app.view = jet.NewSet(jet.NewOSFileSystemLoader("./views"), jet.InDevelopmentMode())
	}

	// init session manager
	app.session = scs.New()
	app.session.Lifetime = 24 * time.Hour
	app.session.Cookie.Persist = true
	app.session.Cookie.Name = app.appName
	app.session.Cookie.Domain = app.server.host
	app.session.Cookie.SameSite = http.SameSiteStrictMode

	if err := app.listenAnServe(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hello, world.")
}
