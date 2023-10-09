package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	// Construct sample payload
	payload := jsonResponse{
		Error:   false,
		Message: "You've contacted the Broker",
	}

	// Marshall payload
	out, err := json.MarshalIndent(payload, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	// Write Response/Headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(out)
}
