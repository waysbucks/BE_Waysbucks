package cartdto

type CreateCart struct {
	UserID        int   `json:"user_id"`
	ProductID     int   `json:"product_id"`
	ToppingID     []int `json:"topping_id"`
	TransactionID []int `json:"transaction_id"`
	QTY           int   `json:"qty"`
	SubTotal      int   `json:"subtotal"`
}

type UpdateCart struct {
	UserID    int `json:"user_id"`
	ProductID int `json:"product_id"`
	ToppingID int `json:"topping_id"`
}

type CartResponse struct {
	UserID    int `json:"user_id"`
	ProductID int `json:"product_id"`
	ToppingID int `json:"topping_id"`
}
