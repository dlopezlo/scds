package main

import (
	"log"
	"net/http"

	"github.com/dlopezlo/scds/server"
)

func main() {
	s := server.CreateNewServer()
	s.MountHandlers()
	const APP_PORT = ":3000"

	err := http.ListenAndServe(APP_PORT, s.Router)
	if err != nil {
		log.Printf("Error starting server: %s\n", err)
	}
	log.Printf("Server ready on port %s", APP_PORT)
}
