package server

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
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
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	s.router.HandleFunc("/v1/health", s.handler.HealthCheck)
	s.router.HandleFunc("/order/{remoteId}", s.handler.Order).Methods("POST")
	s.router.HandleFunc("/remoteId/{remoteId}/remoteOrder/{remoteOrderId}/posOrderStatus", s.handler.UpdateOrderStatus).Methods("PUT")
	s.router.HandleFunc("/menuimport/{remoteId}", s.handler.ImportMenu).Methods("GET")
	s.router.HandleFunc("/catalogimportstatuscallback", s.handler.CatalogImportCallback).Methods("POST")
	s.router.Use(internal.AuthMiddleware)

	// listen and serve
	log.Println("Server listening on: http://localhost:8080")

	srv := &http.Server{
		Handler:      s.router,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("failed to start server: %s", err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)

	return nil
}
