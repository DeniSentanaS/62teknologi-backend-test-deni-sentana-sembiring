package response

type CategoriesResponse struct {
	Id    int    `json:"id"`
	Alias string `json:"alias"`
	Title string `json:"title"`
}
