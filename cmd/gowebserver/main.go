package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hasheddan/gowebserver/pkg/srv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is an API endpoint.")
}

func main() {
	s, err := srv.CreateService(":8080", true)

	if err != nil {
		log.Fatal("Failure to create service.")
	}

	s.Handle("/test", handler)

	log.Fatal(s.Start())
}
