package postgres

import (
	"product-service/internal/product"
	"product-service/internal/product/interfaces"

	"gorm.io/gorm"
)

type ProductRepository struct {
	interfaces.ProductRepositoryInterface
	postgres *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		postgres: db,
	}
}

func (repository *ProductRepository) Create(productData *product.Product) (*product.Product, error) {
	productCreated := repository.postgres.Create(productData)

	if productCreated.Error != nil {
		return nil, productCreated.Error
	}

	return productData, nil
}

func (repository *ProductRepository) GetById(ID string) (*product.Product, error) {

	product := &product.Product{}

	productFound := repository.postgres.First(product, "id = ?", ID)

	if productFound.Error != nil {
		return nil, productFound.Error
	}

	return product, nil
}

func (repository *ProductRepository) GetAll() ([]*product.Product, error) {
	var products []*product.Product

	err := repository.postgres.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (repository *ProductRepository) Delete(ID string) (bool, error) {
	deletedProduct := repository.postgres.Delete(ID)

	if deletedProduct.Error != nil {
		return false, deletedProduct.Error
	}
	return true, nil
}

func (repository *ProductRepository) Uptate(ID string, product *product.Product) (*product.Product, error) {
	productUpdated := repository.postgres.Update(ID, product)

	if productUpdated.Error != nil {
		return nil, productUpdated.Error
	}

	return product, nil
}
