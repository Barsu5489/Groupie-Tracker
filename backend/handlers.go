package backend

import (
	"log"
	"net/http"
	"strconv"
)

func artistsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		artists, err := GetAndUnmarshalArtists()
		if err != nil {
			renderError(w, http.StatusInternalServerError, "Internal server error", "Failed to retrieve artists data")
			log.Printf("Error retrieving artists: %v", err)
			return
		}
		renderTemplate(w, "artists.html", artists)
	} else {
		renderError(w, http.StatusMethodNotAllowed, "Method not allowed", "Only Get requests will be processed in this route")
	}
}

func artistDetailsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		queryParams := r.URL.Query()
		idValue := queryParams.Get("id")
		ID, err := strconv.Atoi(idValue)
		if err != nil {
			renderError(w, http.StatusBadRequest, "Bad Request", "Please request artist information from the links in the cards")
			log.Printf("Error converting id param to int value: %v", err)
			return
		}

		if ID <= 0 || ID > 52 {
			renderError(w, http.StatusNotFound, "Not Found", "Only 52 artists information exists, please request artist information from the links in the cards")
			log.Printf("ID out of range: %d", ID)
			return
		}

		relation, err := GetAndUnmarshalRelation(ID)
		if err != nil {
			renderError(w, http.StatusInternalServerError, "Internal server error", "Failed to retrieve relation data")
			log.Printf("Error retrieving relation data: %v", err)
			return
		}

		renderTemplate(w, "relations.html", relation)
	} else {
		renderError(w, http.StatusMethodNotAllowed, "Method not allowed", "Only Get requests will be processed in this route")
	}
}

func datesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		queryParams := r.URL.Query()
		idValue := queryParams.Get("id")
		ID, err := strconv.Atoi(idValue)
		if err != nil {
			renderError(w, http.StatusBadRequest, "Bad Request", "Please request artist information from the links in the cards")
			log.Printf("Error converting id param to int value: %v", err)
			return
		}

		if ID <= 0 || ID > 52 {
			renderError(w, http.StatusNotFound, "Not Found", "Only 52 artists information exists, please request artist information from the links in the cards")
			log.Printf("ID out of range: %d", ID)
			return
		}

		date, err := GetAndUnmarshalDates(ID)
		if err != nil {
			renderError(w, http.StatusInternalServerError, "Internal server error", "Failed to retrieve dates data")
			log.Printf("Error retrieving relation data: %v", err)
			return
		}

		renderTemplate(w, "dates.html", date)
	} else {
		renderError(w, http.StatusMethodNotAllowed, "Method not allowed", "Only Get requests will be processed in this route")
	}
}

func locationsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		queryParams := r.URL.Query()
		idValue := queryParams.Get("id")
		ID, err := strconv.Atoi(idValue)
		if err != nil {
			renderError(w, http.StatusBadRequest, "Bad Request", "Please request artist information from the links in the cards")
			log.Printf("Error converting id param to int value: %v", err)
			return
		}

		if ID <= 0 || ID > 52 {
			renderError(w, http.StatusNotFound, "Not Found", "Only 52 artists information exists, please request artist information from the links in the cards")
			log.Printf("ID out of range: %d", ID)
			return
		}

		location, err := GetAndUnmarshalLocations(ID)
		if err != nil {
			renderError(w, http.StatusInternalServerError, "Internal server error", "Failed to retrieve locations data")
			log.Printf("Error retrieving relation data: %v", err)
			return
		}

		renderTemplate(w, "locations.html", location)
	} else {
		renderError(w, http.StatusMethodNotAllowed, "Method not allowed", "Only Get requests will be processed in this route")
	}
}
