package endpoint

import (
	"encoding/json"
	"github.com/boinged/api-app-go/model"
	"net/http"
	"go.mongodb.org/mongo-driver/mongo"
)

type MessageEndpoint struct {
	Collection *mongo.Collection
}

func (endpoint *MessageEndpoint) Execute(response http.ResponseWriter, request *http.Request) {
	var result = model.MessageResult{Message: "Hello world!"}
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)
}
