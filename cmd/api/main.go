package main

import (
	"fmt"
	"log"
	"net/http"
)

type Config struct{}

func main() {
	app := Config{}

	// Create server
	srv := &http.Server{
		Addr:    ":80",
		Handler: app.routes(),
	}

	// Start server
	fmt.Printf("** Listening on port 80 **")
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}
