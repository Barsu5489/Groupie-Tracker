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

func GetAndUnmarshalArtists() ([]Artists, error) {
	artists := []Artists{}

	if artistsData != nil {
		err := json.Unmarshal(artistsData, &artists)
		if err != nil {
			return artists, err
		}
		return artists, nil
	}

	jsonData, err := getJSONData("/artists")
	if err != nil {
		return artists, err
	}

	artistsData = jsonData

	err = json.Unmarshal(jsonData, &artists)
	if err != nil {
		return artists, err
	}

	return artists, nil
}

func GetAndUnmarshalLocations(Id int) (Location, error) {
	locations := Locations{}
	location := Location{}

	if locationsData != nil {
		err := json.Unmarshal(locationsData, &locations)
		if err != nil {
			return location, err
		}

		for _, v := range locations.Index {
			if v.ID == Id {
				location = v
			}
		}

		return location, nil
	}

	jsonData, err := getJSONData("/locations")
	if err != nil {
		return location, err
	}

	locationsData = jsonData

	err = json.Unmarshal(jsonData, &locations)
	if err != nil {
		return location, err
	}

	for _, v := range locations.Index {
		if v.ID == Id {
			location = v
		}
	}

	return location, nil
}
