package main

import (
	"net/http"
	"time"

	routes "github.com/simple-me/golang-crud/routes"
)

func main() {
	r := routes.StartGin()
	server := http.Server{
		Addr:           ":8000",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: http.DefaultMaxHeaderBytes,
	}
	server.ListenAndServe()
}
