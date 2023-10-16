package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "8081"

type Config struct{}

func main() {
	app := Config{}

	// Create server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}
	// Start server
	fmt.Printf("** Listening on port %s **\n", webPort)
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}
