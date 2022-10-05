package main

import (
	"github.com/boinged/api-app-go/config"
	"github.com/boinged/api-app-go/endpoint"
	"net/http"
)

func main() {
	http.HandleFunc("/", endpoint.Health)
	http.HandleFunc("/message", endpoint.Message)
	http.ListenAndServe(":"+config.Port, nil)
}
