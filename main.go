package main

import (
	"fmt"
	"net/http"
	"github.com/boinged/api-app-go/config"
)

func main() {
	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprint(response, "Welcome")
	})
	http.ListenAndServe(":" + config.Port, nil)
}
