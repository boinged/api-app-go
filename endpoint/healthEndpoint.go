package endpoint

import (
	"encoding/json"
	"github.com/boinged/api-app-go/model"
	"net/http"
)

type HealthEndpoint struct {
}

func (endpoint *HealthEndpoint) Execute(response http.ResponseWriter, request *http.Request) {
	var result = model.HealthResult{Health: "OK"}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)
}
