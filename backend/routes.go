package backend

import (
	"net/http"
)

// RegisterRoutes returns a serve multiplexer router with all routes registered to their handlers
func RegisterRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("templates"))))
	mux.HandleFunc("/", ArtistsHandler)
	mux.HandleFunc("/locations", LocationsHandler)
	mux.HandleFunc("/dates", DatesHandler)
	mux.HandleFunc("/relation", ArtistDetailsHandler)
	mux.HandleFunc("/api/search", Search)
	return mux
}
