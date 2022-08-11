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

func (h *Handler) UpdateOrderStatus(w http.ResponseWriter, r *http.Request) {

	//get remoteId and remoteOrderId from url
	vars := mux.Vars(r)
	remoteId := vars["remoteId"]
	remoteOrderId := vars["remoteOrderId"]

	// if remoteId is empty, return error
	if remoteId == "" {
		w.WriteHeader(http.StatusBadRequest)
		response := response.CreateAErrorResponse("remote_id is empty")
		w.Write(response)
		return
	}
	// if remoteOrderId is empty, return error
	if remoteOrderId == "" {
		w.WriteHeader(http.StatusBadRequest)
		response := response.CreateAErrorResponse("remote_order_id is empty")
		w.Write(response)
		return
	}

	// convert body to order status
	orderStatus := &requests.OrderStatus{}
	if err := json.NewDecoder(r.Body).Decode(orderStatus); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		// create a response
		response := response.CreateAErrorResponse("invalid request body")
		w.Write(response)
		return
	}

	// log request
	log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
	// log request body json
	log.Printf("%s", orderStatus)
}

// create a menu import endpoint
func (h *Handler) ImportMenu(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// read params from url
	vars := mux.Vars(r)
	remoteId := vars["remoteId"]

	// log remoteId
	log.Printf("update menu import remoteid: %s", remoteId)

	// log request
	log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
}

func (h *Handler) CatalogImportCallback(w http.ResponseWriter, r *http.Request) {

	// get catalogImportCallback from url
	vars := mux.Vars(r)
	catalogImportCallback := vars["catalogImportCallback"]

	//log catalogImportCallback
	log.Printf("catalogImportCallback: %s", catalogImportCallback)

	// parse body to catalogimport
	catalogImport := &requests.CatalogImportRequest{}
	if err := json.NewDecoder(r.Body).Decode(catalogImport); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		// create a response
		response := response.CreateAErrorResponse("invalid request body")
		w.Write(response)
		return
	}

	w.WriteHeader(http.StatusOK)

	// write request body to log in json
	log.Printf("payload: %s", catalogImport)

	// log request
	log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
}
