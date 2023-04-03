package response

type LocationsResponse struct {
	Id        int     `json:"id"`
	Address1  string  `json:"address1"`
	Address2  string  `json:"address2"`
	Address3  string  `json:"address3"`
	City      string  `json:"city"`
	ZipCode   string  `json:"zipCode"`
	State     string  `json:"state"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
