package services

import "product/internal/model/web"

type ProductService interface {
	Create(request web.ProductCreateRequest) (*web.ProductResponse, error)
	Update(request web.ProductUpdateRequest) (*web.ProductResponse, error)
	Delete(productId int) error
	FindById(productId int) (*web.ProductResponse, error)
	FindAll() ([]web.ProductResponse, error)
}
