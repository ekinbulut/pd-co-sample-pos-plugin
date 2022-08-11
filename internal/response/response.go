package response

import (
	"encoding/json"
	"log"
)

type response struct {
	RemoteResponse remoteResponse `json:"remoteResponse"`
}

type remoteResponse struct {
	RemoteOrderId string `json:"remoteOrderId"`
}

func CreateResponse(message string) []byte {
	RemoteResponse := &response{
		RemoteResponse: remoteResponse{
			RemoteOrderId: message,
		},
	}

	response, err := json.Marshal(RemoteResponse)
	if err != nil {
		log.Printf("error marshalling response: %s", err)
	}
	return response
}
