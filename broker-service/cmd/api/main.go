package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = "80"

type Config struct{}

func main() {
	app := Config{}
	log.Printf("Starting broker service on port %s\n", port)

	svr := http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: app.routes(),
	}

	// Starting server
	err := svr.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
