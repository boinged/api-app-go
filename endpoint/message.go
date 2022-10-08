package endpoint

import (
	"encoding/json"
	"github.com/boinged/api-app-go/model"
	"net/http"
)

func Message(response http.ResponseWriter, request *http.Request) {
	var result = model.MessageResult{Message: "Hello world!"}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)
}
