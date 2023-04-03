package request

type CategoriesRequest struct {
	Alias string `json:"alias" binding:"required"`
	Title string `json:"title" binding:"required"`
}
