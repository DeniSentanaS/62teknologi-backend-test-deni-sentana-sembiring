package join

type BusinessesCategories struct {
	BusinessesEntityId int `json:"businessId"`
	CategoriesEntityId int `json:"categoriesId"`
}
type BusinessesLocations struct {
	BusinessesEntityId int `json:"businessId"`
	LocationsEntityId  int `json:"locationsId"`
}
type BusinessesTransactions struct {
	BusinessesEntityId   int `json:"businessId"`
	TransactionsEntityId int `json:"transactionsId"`
}
