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

const APIURL = "https://groupietrackers.herokuapp.com/api"

func getJSONData(url string) (json.RawMessage, error) {
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get json data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received a non-200 response code: %d", resp.StatusCode)
	}

	var jsonString json.RawMessage
	err = json.NewDecoder(resp.Body).Decode(&jsonString)
	if err != nil {
		return nil, fmt.Errorf("failed to decode json data: %w", err)
	}

	return jsonString, nil
}
