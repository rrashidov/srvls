package main

import "net/http"

type application struct {
}

func main() {
	app := &application{}

	srv := &http.Server{
		Handler: app.routes(),
	}

	srv.ListenAndServe()
}
