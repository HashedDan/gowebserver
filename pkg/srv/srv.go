package srv

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

// Service defines the methods the service must implement
type Service interface {
	Start() error
	Stop() error
	Handle() error
}

// Srv is the service that runs our REST server
type Srv struct {
	logging bool
	server  *http.Server
}

// CreateService returns a pointer to a new service instance
func CreateService(address string, log bool) (*Srv, error) {
	httpServ := &http.Server{
		Addr: address,
	}

	return &Srv{
		logging: log,
		server:  httpServ,
	}, nil
}

// Start causes the http server to begin listening for connection
func (s *Srv) Start() error {
	fmt.Printf("Starting server on port %s.\n", s.server.Addr)
	return s.server.ListenAndServe()
}

// Stop closes all connections and shuts down the http server
func (s *Srv) Stop() error {
	fmt.Printf("Shutting down server running on port %s.\n", s.server.Addr)
	return s.server.Shutdown(context.Background())
}

// Handle causes the http server to begin listening for connection
func (s *Srv) Handle(pattern string, handler func(http.ResponseWriter, *http.Request)) error {
	if s.logging {
		fmt.Printf("Adding new server endpoint at %s.\n", pattern)
		http.HandleFunc(pattern, logger(pattern, handler))
		return nil
	}
	http.HandleFunc(pattern, handler)
	return nil
}

func logger(pattern string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Endpoint connection at %s from %s", pattern, r.RemoteAddr)
		next.ServeHTTP(w, r)
	}
}
