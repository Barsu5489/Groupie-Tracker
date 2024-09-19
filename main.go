package main

import (
	"log"
	"net/http"

	"groupie-tracker/backend"
)

// main sets up a simple http server, logs the url and initializes the server to listen and serve on port 8080
func main() {
	server := http.Server{
		Addr:    ":8000",
		Handler: backend.RegisterRoutes(),
	}
	log.Println("server listening on http://localhost:8000")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
