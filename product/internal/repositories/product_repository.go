package repositories

import (
	"gorm.io/gorm"
	"product/internal/model/entity"
)

type ProductRepository interface {
	Save(db *gorm.DB, product *entity.Product) (*entity.Product, error)
	Update(db *gorm.DB, product *entity.Product) (*entity.Product, error)
	Delete(db *gorm.DB, product *entity.Product) error
	FindById(db *gorm.DB, productId int) (*entity.Product, error)
	FindAll(db *gorm.DB) ([]entity.Product, error)
}
