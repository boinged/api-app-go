package endpoint

import (
	"fmt"
	"net/http"
)

func Health(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "OK")
}
