package repositories

import (
	"errors"
	"gorm.io/gorm"
	"product/internal/model/entity"
)

type ProductRepositoryImpl struct{}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

// clear
func (repo ProductRepositoryImpl) Save(db *gorm.DB, product *entity.Product) (*entity.Product, error) {
	err := db.Create(&product).Error
	if err != nil {
		return nil, err
	}

	return product, err
}

// clear
func (repo ProductRepositoryImpl) Update(db *gorm.DB, product *entity.Product) (*entity.Product, error) {
	err := db.Model(&entity.Product{}).Where("id = ?", product.ID).Updates(product).Error
	if err != nil {
		return nil, err
	}

	return product, nil
}

// clear
func (repo ProductRepositoryImpl) Delete(db *gorm.DB, product *entity.Product) error {
	err := db.Delete(&entity.Product{}, "id = ?", product.ID).Error
	if err != nil {
		return err
	}

	return nil
}

// clear
func (repo ProductRepositoryImpl) FindById(db *gorm.DB, productId int) (*entity.Product, error) {
	var product entity.Product

	err := db.Take(&product, "id = ?", productId).Error
	if err != nil {
		return nil, errors.New("data not found")
	}

	return &product, nil
}

// clear
func (repo ProductRepositoryImpl) FindAll(db *gorm.DB) ([]entity.Product, error) {
	var products []entity.Product
	err := db.Find(&products).Error

	if err != nil {
		return nil, err
	}
	return products, nil
}
