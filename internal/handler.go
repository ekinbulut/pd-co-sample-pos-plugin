package internal

import (
	"encoding/json"
	"log"
	"net/http"
	"pos-plugin/internal/requests"
	"pos-plugin/internal/response"

	"github.com/gorilla/mux"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// log request
	log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
}

func (h *Handler) Order(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	remoteId := vars["remoteId"]

	// get auth token from header
	authToken := r.Header.Get("Authorization")
	if authToken == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write(response.CreateAErrorResponse("missing auth token"))
		return
	}

	// if remoteId is empty, return error
	if remoteId == "" {
		response := response.CreateAErrorResponse("remote_id is empty")
		w.Write(response)
		return
	}

	// convert body to order
	order := &requests.Order{}
	if err := json.NewDecoder(r.Body).Decode(order); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		// create a response

		response := response.CreateAErrorResponse("invalid request body")
		w.Write(response)
		return
	}

	response := response.CreateResponse(remoteId)
	w.Write(response)

	// log request
	log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
}
