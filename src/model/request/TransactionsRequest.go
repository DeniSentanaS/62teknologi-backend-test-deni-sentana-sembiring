package request

type TransactionsRequest struct {
	Transaction string `json:"transaction" binding:"required"`
}
