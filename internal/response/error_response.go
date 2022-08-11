package response

import (
	"encoding/json"
	"log"
)

type errorResponse struct {
	Reason  string `json:"reason"`
	Message string `json:"message"`
}

func CreateAErrorResponse(message string) []byte {
	ErrorResponse := &errorResponse{
		Reason:  "FAIL",
		Message: message,
	}

	response, err := json.Marshal(ErrorResponse)
	if err != nil {
		log.Printf("error marshalling response: %s", err)
	}
	return response
}
