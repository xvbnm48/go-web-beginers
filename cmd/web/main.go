package main

import (
	"fmt"
	"log"
	"os"
)

type application struct {
	appName string
	server  server
	debug   bool
	errLog  *log.Logger
	infoLog *log.Logger
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

	if err := app.listenAnServe(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hello, world.")
}
