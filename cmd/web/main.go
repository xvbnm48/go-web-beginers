package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/CloudyKit/jet/v6"
	"github.com/alexedwards/scs/postgresstore"
	"github.com/alexedwards/scs/v2"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"

	_ "github.com/lib/pq"
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

	db2, err := openDB("postgres://fariz:fariz@localhost/hackernews?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	defer db2.Close()

	// init upperdb
	upper, err := postgresql.New(db2)
	if err != nil {
		log.Fatal(err)
	}

	defer func(upper db.Session) {
		err := upper.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(upper)

	// init application
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
	app.session.Store = postgresstore.New(db2)

	if err := app.listenAnServe(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hello, world.")
}

//OpenDB: open database connection
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
