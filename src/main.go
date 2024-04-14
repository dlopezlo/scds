package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
  fmt.Printf("Starting web server...")
	r := chi.NewRouter()
	r.Use(middleware.Timeout(time.Second * 3))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Welcome to my TFG"))
		if err != nil {
			fmt.Println(err)
		}
	})
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		panic(err)
	}

}

