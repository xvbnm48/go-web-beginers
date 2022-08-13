package main

import (
	"fmt"
	"net/http"
	"time"
)

func (a *application) listenAnServe() error {
	host := fmt.Sprintf("%s:%s", a.server.host, a.server.port)
	srv := http.Server{
		Handler:     a.routes(),
		Addr:        host,
		ReadTimeout: 10 * time.Second,
	}

	a.infoLog.Printf("Listening on %s\n", host)

	return srv.ListenAndServe()
}
