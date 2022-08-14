package models

import "time"

type Product struct {
	ID        int       `json:"id"`
	Title     string    `json:"title" gorm:"type: varchar(255)"`
	Price     int       `json:"price" gorm:"type: int"`
	Image     string    `json:"image" gorm:"type: varchar(255)"`
	Topping   []Topping `json:"topping" gorm:"many2many:product_topping"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type ProductTransaction struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Price int    `json:"price"`
	Image string `json:"image"`
}

func (ProductTransaction) TableName() string {
	return "products"
}
