package web

type ProductUpdateRequest struct {
	Id       int    `validate:"required"`
	Name     string `form:"name" validate:"required"`
	Quantity int    `form:"quantity" validate:"required"`
	Price    int    `form:"price" validate:"required"`
	Image    string `form:"image"`
}
