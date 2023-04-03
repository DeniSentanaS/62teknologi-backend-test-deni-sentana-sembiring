package request

type LocationsRequest struct {
	Address1  string  `json:"address1" binding:"required"`
	Address2  string  `json:"address2"`
	Address3  string  `json:"address3"`
	City      string  `json:"city" binding:"required"`
	ZipCode   string  `json:"zipCode" binding:"required"`
	State     string  `json:"state" binding:"required"`
	Latitude  float64 `json:"latitude" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
}
