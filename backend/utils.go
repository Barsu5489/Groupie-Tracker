package backend

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"
	"time"
)

var client = &http.Client{
	Timeout: 15 * time.Second,
}

// fetched json cache
var (
	artistsData   json.RawMessage
	locationsData json.RawMessage
	datesData     json.RawMessage
	relationData  json.RawMessage
)

const APIURL = "https://groupietrackers.herokuapp.com/api"

// getJSONData fetches JSON data from the API and returns it as a RawMessage
func GetJSONData(baseURL, endpoint string, client *http.Client) (json.RawMessage, error) {
	resp, err := client.Get(baseURL + endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to get %s json data: %w", endpoint, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received a non-200 response code for %s: %d", endpoint, resp.StatusCode)
	}

	var jsonString json.RawMessage
	err = json.NewDecoder(resp.Body).Decode(&jsonString)
	if err != nil {
		return nil, fmt.Errorf("failed to decode %s json data: %w", endpoint, err)
	}

	return jsonString, nil
}

// renderTemplate renders the correct html templates
func renderTemplate(w http.ResponseWriter, templateName string, data interface{}) {
	tempFile := filepath.Join("templates", templateName)
	temp, err := template.ParseFiles(tempFile)
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Internal Server Error", "Failed to load template")
		log.Printf("Error parsing template file: %v", err)
		return
	}

	err = temp.Execute(w, data)
	if err != nil {
		renderError(w, http.StatusInternalServerError, "Internal Server Error", "Failed to render template")
		log.Printf("Error executing template: %v", err)
	}
}

// unmarshalData is a helper function to unmarshal cached JSON data or fetch it if not available
func unmarshalData(cache *json.RawMessage, endpoint string, out interface{}) error {
	if *cache != nil {
		return json.Unmarshal(*cache, out)
	}

	jsonData, err := GetJSONData(APIURL, endpoint, client)
	if err != nil {
		return err
	}

	*cache = jsonData
	return json.Unmarshal(jsonData, out)
}

// GetAndUnmarshalArtists returns a list of artists by fetching or using cached data
func GetAndUnmarshalArtists() ([]Artists, error) {
	artists := []Artists{}
	err := unmarshalData(&artistsData, "/artists", &artists)
	return artists, err
}

// GetAndUnmarshalLocations returns a specific location by its ID
func GetAndUnmarshalLocations(ID int) (Location, error) {
	locations := Locations{}
	err := unmarshalData(&locationsData, "/locations", &locations)
	if err != nil {
		return Location{}, err
	}

	for _, v := range locations.Index {
		if v.ID == ID {
			return v, nil
		}
	}
	return Location{}, fmt.Errorf("location with ID %d not found", ID)
}
func GetAndUnmarshalLocation() (Locations, error) {
	locations := Locations{}
	err := unmarshalData(&locationsData, "/locations", &locations)
	if err != nil {
		return Locations{}, err
	}

	return locations, nil
}

// GetAndUnmarshalDates returns specific dates by their ID
func GetAndUnmarshalDates(ID int) (Date, error) {
	dates := Dates{}
	err := unmarshalData(&datesData, "/dates", &dates)
	if err != nil {
		return Date{}, err
	}

	for _, v := range dates.Index {
		if v.ID == ID {
			return v, nil
		}
	}
	return Date{}, fmt.Errorf("date with ID %d not found", ID)
}
func GetAndUnmarshalDate() (Dates, error) {
	dates := Dates{}
	err := unmarshalData(&datesData, "/dates", &dates)
	if err != nil {
		return Dates{}, err
	}
	return dates, nil
}

// GetAndUnmarshalRelation returns artist details by their ID
func GetAndUnmarshalRelation(ID int) (ArtistDetails, error) {
	relation := Relation{}
	err := unmarshalData(&relationData, "/relation", &relation)
	if err != nil {
		return ArtistDetails{}, err
	}

	for _, v := range relation.Index {
		if v.ID == ID {
			return v, nil
		}
	}
	return ArtistDetails{}, fmt.Errorf("relation with ID %d not found", ID)
}

// renderError renders http errors status codes with their descriptions
func renderError(w http.ResponseWriter, status int, errorMessage string, information string) {
	w.WriteHeader(status)
	tempFile := filepath.Join("templates", "errors.html")
	t, err := template.ParseFiles(tempFile)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	data := struct {
		Status int
		Error  string
		Info   string
	}{
		Status: status,
		Error:  errorMessage,
		Info:   information,
	}
	t.Execute(w, data)
}
