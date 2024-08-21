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
			http.Error(w, "Failed to retrieve artists data", http.StatusInternalServerError)
			log.Printf("Error retrieving artists: %v", err)
			return
		}
		renderTemplate(w, "artists.html", artists)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func artistDetailsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		queryParams := r.URL.Query()
		idValue := queryParams.Get("id")
		ID, err := strconv.Atoi(idValue)

		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			log.Printf("Error converting id param to int value: %v", err)
			return
		}

		if ID <= 0 || ID > 52 {
			http.Error(w, "ID out of range", http.StatusBadRequest)
			log.Printf("ID out of range: %d", ID)
			return
		}

		relation, err := GetAndUnmarshalRelation(ID)
		if err != nil {
			http.Error(w, "Failed to retrieve relation data", http.StatusInternalServerError)
			log.Printf("Error retrieving relation data: %v", err)
			return
		}

		renderTemplate(w, "relations.html", relation)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
func datesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		queryParams := r.URL.Query()
		idValue := queryParams.Get("id")
		ID, err := strconv.Atoi(idValue)

		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			log.Printf("Error converting id param to int value: %v", err)
			return
		}

		if ID <= 0 || ID > 52 {
			http.Error(w, "ID out of range", http.StatusBadRequest)
			log.Printf("ID out of range: %d", ID)
			return
		}

		date, err := GetAndUnmarshalDates(ID)
		if err != nil {
			http.Error(w, "Failed to retrieve relation data", http.StatusInternalServerError)
			log.Printf("Error retrieving relation data: %v", err)
			return
		}

		renderTemplate(w, "dates.html", date)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
func locationsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		queryParams := r.URL.Query()
		idValue := queryParams.Get("id")
		ID, err := strconv.Atoi(idValue)

		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			log.Printf("Error converting id param to int value: %v", err)
			return
		}

		if ID <= 0 || ID > 52 {
			http.Error(w, "ID out of range", http.StatusBadRequest)
			log.Printf("ID out of range: %d", ID)
			return
		}

		location, err := GetAndUnmarshalLocations(ID)
		if err != nil {
			http.Error(w, "Failed to retrieve relation data", http.StatusInternalServerError)
			log.Printf("Error retrieving relation data: %v", err)
			return
		}

		renderTemplate(w, "locations.html", location)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
