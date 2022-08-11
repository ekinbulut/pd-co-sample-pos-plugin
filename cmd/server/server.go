package server

import (
	"log"
	"net/http"
	"pos-plugin/internal"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	Version string
	router  *mux.Router
	handler *internal.Handler
}

func NewServer() *Server {
	return &Server{
		Version: "1.0.0",
		router:  mux.NewRouter(),
		handler: internal.NewHandler(),
	}
}

func (s *Server) Start() error {
	// register handlers

	s.router.HandleFunc("/v1/health", s.handler.HealthCheck)
	s.router.HandleFunc("/order/{remoteId}", s.handler.Order).Methods("POST")
	s.router.HandleFunc("/remoteId/{remoteId}/remoteOrder/{remoteOrderId}/posOrderStatus", s.handler.UpdateOrderStatus).Methods("PUT")
	s.router.HandleFunc("/menuimport/{remoteId}", s.handler.ImportMenu).Methods("GET")
	s.router.HandleFunc("/{catalogImportCallback}", s.handler.CatalogImportCallback).Methods("POST")

	// listen and serve
	log.Println("Server listening on: http://localhost:8080")

	srv := &http.Server{
		Handler:      s.router,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
