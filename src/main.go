package main

import (
	"log"
	"net/http"

	"github.com/dlopezlo/scds/server"
)

const APP_PORT = ":3000"

func main() {

	s := server.CreateNewServer()
	s.MountHandlers()

	log.Printf("Starting server on port %s", APP_PORT)
	err := http.ListenAndServe(APP_PORT, s.Router)
	if err != nil {
		log.Printf("Error starting server: %s\n", err)
	}
}
