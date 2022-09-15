package interfaces

import "product-service/internal/product"

type ProductServiceInterface interface {
	Create(product *product.Product) (*product.Product, error)
	GetById(ID string) (*product.Product, error)
	GetAll() ([]*product.Product, error)
	Delete(ID string) (bool, error)
	Uptate(ID string, product *product.Product) (*product.Product, error)
}
