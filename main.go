package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprint(response, "Welcome")
	})
	http.ListenAndServe(":8080", nil)
}
