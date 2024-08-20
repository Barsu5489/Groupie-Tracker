package main

import (
	"log"
	"net/http"

	"groupie-tracker/backend"
)

func main() {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: backend.RegisterRoutes(),
	}
	log.Println("server listening on http://localhost:8080")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
