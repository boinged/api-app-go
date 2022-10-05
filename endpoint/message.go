package endpoint

import (
	"encoding/json"
	"net/http"
)

func Message(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode("Hello world!")
}
