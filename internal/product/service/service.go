package service

import (
	"product-service/internal/product"
	"product-service/internal/product/repository"
)

type ProductServiceInterface interface {
	Create(product *product.Product) (*product.Product, error)
	GetById(ID string) (*product.Product, error)
	GetAll() ([]*product.Product, error)
	Delete(ID string) (bool, error)
	Uptate(ID string, product *product.Product) (*product.Product, error)
}

func NewProductService(repository repository.ProductRepositoryInterface) ProductServiceInterface {
	return &ProductService{
		Repository: repository,
	}
}

type ProductService struct {
	Repository repository.ProductRepositoryInterface
}

func (productService *ProductService) Create(product *product.Product) (*product.Product, error) {
	createdProduct, err := productService.Repository.Create(product)
	if err != nil {
		return nil, err
	}
	return createdProduct, nil
}

func (productService *ProductService) GetById(ID string) (*product.Product, error) {
	product, err := productService.Repository.GetById(ID)
	if err != nil {
		return nil, err
	}
	return product, nil
}
func (productService *ProductService) GetAll() ([]*product.Product, error) {
	products, err := productService.Repository.GetAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}
func (productService *ProductService) Delete(ID string) (bool, error) {
	isDeleted, err := productService.Repository.Delete(ID)
	if err != nil {
		return false, err
	}
	return isDeleted, nil
}
func (productService *ProductService) Uptate(ID string, product *product.Product) (*product.Product, error) {
	productUpdated, err := productService.Repository.Uptate(ID, product)
	if err != nil {
		return nil, err
	}
	return productUpdated, nil
}
