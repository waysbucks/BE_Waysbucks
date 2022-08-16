package models

type Cart struct {
	ID            int                `json:"id" gorm:"primary_key:auto_increment"`
	QTY           int                `json:"qty"`
	SubTotal      int                `json:"subtotal"`
	ProductID     int                `json:"product_id"`
	Product       ProductTransaction `json:"product"`
	ToppingID     []int              `json:"topping_id" gorm:"-"`
	Topping       []Topping          `json:"topping" gorm:"many2many:cart_toppings"`
	TransactionID int                `json:"transaction_id"`
	Transaction   Transaction        `json:"transaction"`
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
