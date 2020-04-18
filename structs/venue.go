package structs

type Venue struct {
	Location       Location `json:"location"`
	Title          string   `json:"title"`
	Address        string   `json:"address"`
	FoursquareId   string   `json:"foursquare_id"`
	FoursquareType string   `json:"foursquare_type"`
}
