package transactiondto

type CreateTransaction struct {
	UserID    int `json:"user_id" form:"user_id" validate:"required"`
	ProductID int `json:"product_id" form:"product_id" validate:"required"`
	ToppingID int `json:"topping_id" form:"topping_id" validate:"required"`
}

type UpdateTransaction struct {
	UserID    int `json:"user_id" form:"user_id"`
	ProductID int `json:"product_id" form:"product_id"`
	ToppingID int `json:"topping_id" form:"topping_id"`
}

// type TransactionResponse struct {
// 	UserID    int `json:"user_id" form:"user_id"`
// 	ProductID int `json:"product_id" form:"product_id"`
// 	ToppingID int `json:"topping_id" form:"topping_id"`
// }
