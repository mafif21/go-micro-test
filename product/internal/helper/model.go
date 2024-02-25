package helper

import (
	"product/internal/model/entity"
	"product/internal/model/web"
)

func ToCategoryResponse(product entity.Product) web.ProductResponse {
	return web.ProductResponse{
		Id:         product.ID,
		Name:       product.Name,
		Quantity:   product.Quantity,
		Price:      product.Price,
		Image:      product.Image,
		CreatedAt:  product.CreatedAt,
		Updated_at: product.UpdatedAt,
	}
}
