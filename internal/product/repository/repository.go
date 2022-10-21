package repository

import "product-service/internal/product"

type ProductRepositoryInterface interface {
	Create(product *product.Product) (*product.Product, error)
	GetById(ID string) (*product.Product, error)
	GetAll() ([]*product.Product, error)
	Delete(ID string) (bool, error)
	Uptate(ID string, product *product.Product) (*product.Product, error)
}

type ProductRepository struct {
	Persister ProductRepositoryInterface
}

func NewProductRepository(persister ProductRepositoryInterface) ProductRepositoryInterface {
	return &ProductRepository{
		Persister: persister,
	}
}

func (pr *ProductRepository) Create(product *product.Product) (*product.Product, error) {
	return pr.Persister.Create(product)
}

func (pr *ProductRepository) GetById(ID string) (*product.Product, error) {
	return pr.Persister.GetById(ID)
}
func (pr *ProductRepository) GetAll() ([]*product.Product, error) {
	return pr.Persister.GetAll()
}
func (pr *ProductRepository) Delete(ID string) (bool, error) {
	return pr.Persister.Delete(ID)
}
func (pr *ProductRepository) Uptate(ID string, product *product.Product) (*product.Product, error) {
	return pr.Persister.Uptate(ID, product)
}

