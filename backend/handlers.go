package backend

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// artistHandler handles requests to the home page and serves artist cards
func ArtistsHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		renderError(w, http.StatusNotFound, "Not Found", "The requested URL was not found on this server")
		return
	} else if r.Method != http.MethodGet {
		// Handle only GET requests
		renderError(w, http.StatusMethodNotAllowed, "Method Not Allowed", "Only GET requests are allowed on this route")
		return
	}
	artists, err := GetAndUnmarshalArtists()
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Internal server error", "Failed to retrieve artists data")
		log.Printf("Error retrieving artists: %v", err)
		return
	}
	renderTemplate(w, "artists.html", artists)
}

func DetailsHanlder(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/details" {
		renderError(w, http.StatusNotFound, "Not Found", "The requested URL was not found on this server")
		return
	} else if r.Method != http.MethodGet {
		// Handle only GET requests
		renderError(w, http.StatusMethodNotAllowed, "Method Not Allowed", "Only GET requests are allowed on this route")
		return
	}
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
	artists, err := GetAndUnmarshalArtists()
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Internal server error", "Failed to retrieve artists data")
		log.Printf("Error retrieving artists: %v", err)
		return
	}

	var selectedArtist Artists // Define Artist struct as per your model
	for _, artist := range artists {
		if artist.ID == ID { // Assuming each artist has an ID field
			selectedArtist = artist
			break
		}
	}

	relation, err := GetAndUnmarshalRelation(ID)
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Internal server error", "Failed to retrieve relation data")
		log.Printf("Error retrieving relation data: %v", err)
		return
	}

	date, err := GetAndUnmarshalDates(ID)
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Internal server error", "Failed to retrieve dates data")
		log.Printf("Error retrieving relation data: %v", err)
		return
	}

	location, err := GetAndUnmarshalLocations(ID)
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Internal server error", "Failed to retrieve locations data")
		log.Printf("Error retrieving relation data: %v", err)
		return
	}

	details := BandDetails{
		Artist: selectedArtist,
		Relation: relation,
		Dates: date,
		Locations: location,
	}
	
	renderTemplate(w, "details.html", details)
}

// artistDetailsHandler handles requests for the artist relation data
func ArtistDetailsHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/relation" {
		renderError(w, http.StatusNotFound, "Not Found", "The requested URL was not found on this server")
		return
	} else if r.Method != http.MethodGet {
		// Handle only GET requests
		renderError(w, http.StatusMethodNotAllowed, "Method Not Allowed", "Only GET requests are allowed on this route")
		return
	}
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
}

// dates Handler handles requests for the dates data
func DatesHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/dates" {
		renderError(w, http.StatusNotFound, "Not Found", "The requested URL was not found on this server")
		return
	} else if r.Method != http.MethodGet {
		// Handle only GET requests
		renderError(w, http.StatusMethodNotAllowed, "Method Not Allowed", "Only GET requests are allowed on this route")
		return
	}
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
}

// locations Handler handles requests for the locations data
func LocationsHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/locations" {
		renderError(w, http.StatusNotFound, "Not Found", "The requested URL was not found on this server")
		return
	} else if r.Method != http.MethodGet {
		// Handle only GET requests
		renderError(w, http.StatusMethodNotAllowed, "Method Not Allowed", "Only GET requests are allowed on this route")
		return
	}
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
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := strings.ToLower(r.URL.Query().Get("q"))

	artists, err := GetAndUnmarshalArtists()
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Internal server error", "Failed to retrieve artists data")
		log.Printf("Error retrieving artists: %v", err)
		return
	}

	suggestions := []Suggestion{}

	for _, v := range artists {
		if strings.Contains(strings.ToLower(v.Name), query) {
			suggestions = append(suggestions, Suggestion{
				ID:   v.ID,
				Text: fmt.Sprintf("%v - artists/band", v.Name),
			})
		}

		for _, mem := range v.Members {
			if strings.Contains(strings.ToLower(mem), query) {
				suggestions = append(suggestions, Suggestion{
					ID:   v.ID,
					Text: fmt.Sprintf("%v - Member of %v", mem, v.Name),
				})
			}
		}

		if strings.Contains(strings.ToLower(v.FirstAlbum), query) {
			suggestions = append(suggestions, Suggestion{
				ID:   v.ID,
				Text: fmt.Sprintf("%v produced first album on %v", v.Name, v.FirstAlbum),
			})
		}

		if strings.Contains(strconv.Itoa(v.CreationDate), query) {
			suggestions = append(suggestions, Suggestion{
				ID:   v.ID,
				Text: fmt.Sprintf("%v created on %v", v.Name, v.CreationDate),
			})
		}
	}

	location, err := GetAndUnmarshalLocation()
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Internal server error", "Failed to retrieve location data")
		log.Printf("Error retrieving location: %v", err)
		return
	}

	for _, loc := range location.Index {
		for _, l := range loc.Locations {
			if strings.Contains(strings.ToLower(l), query) {
				for _, v := range artists {
					if v.ID == loc.ID {
						suggestions = append(suggestions, Suggestion{
							ID:   v.ID,
							Text: fmt.Sprintf("%v performing at %v", v.Name, l),
						})
					}
				}
			}
		}
	}

	dates, err := GetAndUnmarshalDate()
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Internal server error", "Failed to retrieve dates data")
		log.Printf("Error retrieving dates: %v", err)
		return
	}

	for _, date := range dates.Index {
		for _, d := range date.Dates {
			if strings.Contains(strings.ToLower(d), query) {
				for _, v := range artists {
					if v.ID == date.ID {
						suggestions = append(suggestions, Suggestion{
							ID:   v.ID,
							Text: fmt.Sprintf("%v will be performing on %v", v.Name, d),
						})
					}
				}
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(suggestions)
}

// FilterArtists returns artist data for a list of IDs
func FilterArtistsHandler(w http.ResponseWriter, r *http.Request) {
    idsQuery := r.URL.Query()["ids"] // Retrieve 'ids' query parameter as a slice
    if len(idsQuery) == 0 {
        http.Error(w, "No IDs provided", http.StatusBadRequest)
        return
    }

	artists, err := GetAndUnmarshalArtists()
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Internal server error", "Failed to retrieve artists data")
		log.Printf("Error retrieving artists: %v", err)
		return
	}

    var matchingArtists []Artists
    // Assuming 'artists' is a slice of all available artists
    for _, artist := range artists {
        for _, id := range idsQuery {
            if strconv.Itoa(artist.ID) == id {
                matchingArtists = append(matchingArtists, artist)
                break
            }
        }
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(matchingArtists)
}

