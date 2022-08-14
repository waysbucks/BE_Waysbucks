package models

import "time"

type Transaction struct {
	ID        int                `json:"id"`
	UserID    int                `json:"user_id"`
	ProductID int                `json:"product_id"`
	ToppingID int                `json:"topping_id"`
	User      UserProfile        `json:"user"`
	Product   ProductTransaction `json:"product"`
	Topping   ToppingTransaction `json:"topping"`
	CreatedAt time.Time          `json:"-"`
	UpdatedAt time.Time          `json:"-"`
}
