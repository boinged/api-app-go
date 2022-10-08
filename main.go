package main

import (
	"fmt"
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
	fmt.Println(database.Name())

	http.HandleFunc("/", endpoint.Health)
	http.HandleFunc("/message", endpoint.Message)
	http.ListenAndServe(":"+config.Port, nil)
}
