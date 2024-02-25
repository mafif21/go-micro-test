package entity

import "time"

type Product struct {
	ID        int       `gorm:"primaryKey;column:id;autoIncrement"`
	Name      string    `gorm:"column:name"`
	Quantity  int       `gorm:"column:quantity"`
	Price     int       `gorm:"column:price"`
	Image     string    `gorm:"column:image"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (p *Product) TableName() string {
	return "products"
}
