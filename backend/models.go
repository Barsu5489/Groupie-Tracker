package backend

// Artists struct for unmarshalling artists data
type Artists struct {
	ID              int      `json:"id"`
	ImageURL        string   `json:"image"`
	Name            string   `json:"name"`
	Members         []string `json:"members"`
	CreationDate    int      `json:"creationDate"`
	FirstAlbum      string   `json:"firstAlbum"`
	LocationsURL    string   `json:"locations"`
	ConcertDatesURL string   `json:"concertDates"`
	RelationsURL    string   `json:"relations"`
}

// Nested Locations structs for unmarshalling locations data
type Locations struct {
	Index []Location `json:"index"`
}

type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

// Nested Dates structs for unmarshalling dates data
type Dates struct {
	Index []Date `json:"index"`
}

type Date struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

// Nested Relation structs for unmarshalling relation data
type Relation struct {
	Index []ArtistDetails `json:"index"`
}

type ArtistDetails struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type BandDetails struct {
	Artist    Artists  // Replace with your actual artist struct type
	Relation  ArtistDetails // Replace with your actual relation struct type
	Dates     Date // Replace with your actual date struct type
	Locations Location // Replace with your actual location struct type
}

type Suggestion struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}