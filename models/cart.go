package models

type Cart struct {
	ID        int                `json:"id"`
	UserID    int                `json:"user_id"`
	User      UserProfile        `json:"user"`
	ProductID int                `json:"product_id"`
	Product   ProductTransaction `json:"product"`
	ToppingID int                `json:"topping_id"`
	Topping   ToppingTransaction `json:"topping"`
}

type UserCart struct {
	ID        int                `json:"id"`
	UserID    int                `json:"user_id"`
	ProductID int                `json:"product_id"`
	ToppingID int                `json:"topping_id"`
	Product   ProductTransaction `json:"product"`
	Topping   ToppingTransaction `json:"topping"`
}

func (UserCart) TableName() string {
	return "carts"
}
