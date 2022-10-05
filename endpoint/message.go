package endpoint

import (
	"fmt"
	"net/http"
)

func Message(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "Hello world!")
}
