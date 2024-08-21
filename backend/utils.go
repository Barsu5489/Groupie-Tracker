package backend

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var client = &http.Client{
	Timeout: 10 * time.Second,
}

var (
	artistsData   json.RawMessage
	locationsData json.RawMessage
	datesData     json.RawMessage
	relationData  json.RawMessage
)

const APIURL = "https://groupietrackers.herokuapp.com/api"

// getJSONData fetches JSON data from the API and returns it as a RawMessage
func getJSONData(endpoint string) (json.RawMessage, error) {
	resp, err := client.Get(APIURL + endpoint)
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

// unmarshalData is a helper function to unmarshal cached JSON data or fetch it if not available
func unmarshalData(cache *json.RawMessage, endpoint string, out interface{}) error {
	if *cache != nil {
		return json.Unmarshal(*cache, out)
	}

	jsonData, err := getJSONData(endpoint)
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
