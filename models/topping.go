package models

import "time"

type Topping struct {
	ID        int       `json:"id"`
	Title     string    `json:"title" gorm:"type: varchar(255)"`
	Price     int       `json:"price" gorm:"type: int"`
	Image     string    `json:"image" gorm:"type: varchar(255)"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type ToppingTransaction struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Price int    `json:"price"`
	Image string `json:"image"`
}

func (ToppingTransaction) TableName() string {
	return "toppings"
}
