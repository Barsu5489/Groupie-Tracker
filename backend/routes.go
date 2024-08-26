package backend

import (
	"net/http"
)

func RegisterRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("templates"))))
	mux.HandleFunc("/", artistsHandler)
	mux.HandleFunc("/locations", locationsHandler)
	mux.HandleFunc("/dates", datesHandler)
	mux.HandleFunc("/relation", artistDetailsHandler)
	return mux
}
