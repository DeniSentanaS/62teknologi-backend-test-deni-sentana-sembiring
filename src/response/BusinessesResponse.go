package response

type BusinessesResponse struct {
	Id           int                         `json:"id"`
	Alias        string                      `json:"alias"`
	Name         string                      `json:"name"`
	ImageUrl     string                      `json:"imageUrl"`
	IsClosed     bool                        `json:"isClosed"`
	Url          string                      `json:"url"`
	ReviewCount  int                         `json:"reviewCount"`
	Rating       float64                     `json:"rating"`
	Phone        string                      `json:"phone"`
	DisplayPhone string                      `json:"displayName"`
	Distance     float64                     `json:"distance"`
	Categories   []CategoriesResponse        `json:"categories"`
	Locations    []BusinessLocationsResponse `json:"locations"`
	Transactions []TransactionsResponse      `json:"transactions"`
}
