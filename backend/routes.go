package backend

import (
	"net/http"
)

func RegisterRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	mux.HandleFunc("/", ArtistsHandler)
	mux.HandleFunc("/locations", LocationsHandler)
	mux.HandleFunc("/dates", DatesHandler)
	mux.HandleFunc("/relation", ArtistDetailsHandler)
	mux.HandleFunc("/details", DetailsHanlder)
	mux.HandleFunc("/api/search", SearchHandler)
	mux.HandleFunc("/api/artists/filter", FilterArtistsHandler)
	return mux
}
