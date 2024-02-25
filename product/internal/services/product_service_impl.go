package services

import (
	"fmt"
	"gorm.io/gorm"
	"os"
	"path"
	"product/internal/config"
	"product/internal/exception"
	"product/internal/helper"
	"product/internal/model/entity"
	"product/internal/model/web"
	"product/internal/repositories"
	"time"
)

type ProductServiceImpl struct {
	DB         *gorm.DB
	Repository repositories.ProductRepository
}

func NewProductService(db *gorm.DB, repository repositories.ProductRepository) ProductService {
	return &ProductServiceImpl{
		DB:         db,
		Repository: repository,
	}
}

// clear
func (service ProductServiceImpl) Create(request web.ProductCreateRequest) (*web.ProductResponse, error) {
	db := config.OpenConnection()

	imageUrl := fmt.Sprintf("/public/product/%s", request.Image)
	request.Image = imageUrl

	newProduct := entity.Product{
		Name:     request.Name,
		Quantity: request.Quantity,
		Price:    request.Price,
		Image:    request.Image,
	}

	fmt.Println(newProduct)

	save, err := service.Repository.Save(db, &newProduct)
	if err != nil {
		return nil, &exception.ErrorMessage{Message: "cant save the product"}
	}

	response := helper.ToCategoryResponse(*save)
	return &response, nil
}

// clear
func (service ProductServiceImpl) Update(request web.ProductUpdateRequest) (*web.ProductResponse, error) {
	db := config.OpenConnection()

	productFounded, err := service.Repository.FindById(db, request.Id)
	if err != nil {
		return nil, &exception.ErrorMessage{Message: err.Error()}
	}

	if request.Image == "" {
		request.Image = productFounded.Image
	}

	if productFounded.Image != request.Image {
		_ = os.Remove("./public/product/" + path.Base(productFounded.Image))
		productFounded.Image = request.Image
	} else {
		productFounded.Image = productFounded.Image
	}

	productFounded.Name = request.Name
	productFounded.Quantity = request.Quantity
	productFounded.Price = request.Price
	productFounded.UpdatedAt = time.Now()

	productUpdate, err := service.Repository.Update(db, productFounded)
	if err != nil {
		return nil, &exception.ErrorMessage{Message: "cant save the product"}
	}

	response := helper.ToCategoryResponse(*productUpdate)
	return &response, nil
}

// clear
func (service ProductServiceImpl) Delete(productId int) error {
	db := config.OpenConnection()
	productFounded, err := service.Repository.FindById(db, productId)
	if err != nil {
		return &exception.ErrorMessage{Message: err.Error()}
	}

	err = service.Repository.Delete(db, productFounded)
	if err != nil {
		return &exception.ErrorMessage{Message: "cant delete the product"}
	}

	_ = os.Remove("./public/product/" + path.Base(productFounded.Image))

	return nil
}

// clear
func (service ProductServiceImpl) FindById(productId int) (*web.ProductResponse, error) {
	db := config.OpenConnection()

	productFounded, err := service.Repository.FindById(db, productId)
	if err != nil {
		return nil, &exception.ErrorMessage{Message: err.Error()}
	}

	response := helper.ToCategoryResponse(*productFounded)
	return &response, nil
}

// clear
func (service ProductServiceImpl) FindAll() ([]web.ProductResponse, error) {
	var allProducts []web.ProductResponse

	db := config.OpenConnection()
	products, err := service.Repository.FindAll(db)
	if err != nil {
		return nil, &exception.ErrorMessage{Message: "cant get all product"}
	}

	for _, product := range products {
		allProducts = append(allProducts, helper.ToCategoryResponse(product))
	}

	return allProducts, nil
}
