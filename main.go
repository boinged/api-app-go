package main

import (
	"net/http"

	"github.com/boinged/api-app-go/config"
	"github.com/boinged/api-app-go/database"
	"github.com/boinged/api-app-go/endpoint"
)

func main() {
	if config.DatabaseURI == "" {
		panic("missing database uri")
	}

	database := database.Connect(config.DatabaseURI)
	collection := database.Collection("message")

	healthEndpoint := endpoint.HealthEndpoint{}
	http.HandleFunc("/", healthEndpoint.Execute)

	messageEndpoint := endpoint.MessageEndpoint{Collection: collection}
	http.HandleFunc("/message", messageEndpoint.Execute)

	http.ListenAndServe(":"+config.Port, nil)
}
