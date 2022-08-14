package productdto

type CreateProduct struct {
	Title string `json:"title" form:"title" validate:"required"`
	Price int    `json:"price" form:"price" gorm:"type: int" validate:"required"`
	Image string `json:"image" form:"id" validate:"required"`
}

type UpdateProduct struct {
	Title string `json:"title" form:"title"`
	Price int    `json:"price" form:"price"`
	Image string `json:"image" form:"id"`
}

type ProductResponse struct {
	Title string `json:"title" form:"title"`
	Price int    `json:"price" form:"price"`
	Image string `json:"image" form:"image"`
}
