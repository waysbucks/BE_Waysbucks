package cartdto

type CreateCart struct {
	ID        int   `json:"id"`
	UserID    int   `json:"user_id"`
	ProductID int   `json:"product_id"`
	ToppingID []int `json:"topping_id"`
	QTY       int   `json:"qty"`
	SubTotal  int   `json:"subtotal"`
}

type UpdateCart struct {
	UserID    int `json:"user_id"`
	ProductID int `json:"product_id"`
	ToppingID int `json:"topping_id"`
}

type CartResponse struct {
	ID       int `json:"id"`
	QTY      int `json:"qty"`
	SubTotal int `json:"subtotal"`
}
