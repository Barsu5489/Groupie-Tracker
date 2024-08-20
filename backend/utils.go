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
	relationsData json.RawMessage
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

	err = json.Unmarshal(jsonData, &artists)
	if err != nil {
		return artists, err
	}

	return artists, nil
}
