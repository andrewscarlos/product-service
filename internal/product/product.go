package product

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	Value     float64   `json:"value"`
	Total     float64   `json:"total"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Order struct {
	ID        uuid.UUID `json:"id"`
	Product   Product
	ProductID uint `sql:"product_id"`
}
