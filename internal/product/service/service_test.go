package service_test

import (
	"product-service/internal/product"
	"product-service/internal/product/repository"
	"product-service/internal/product/service"
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestProductService(test *testing.T) {
	productRepositoryMock := &repository.Mock{}
	productService := service.NewProductService(productRepositoryMock)

	test.Run("Create", func(test *testing.T) {
		test.Run("It should create a new product", func(test *testing.T) {
			newProduct := new(product.Product)
			newProduct.ID = uuid.NewV4()
			newProduct.Name = "TV"
			newProduct.Type = "foo"
			newProduct.Value = 100

			productRepositoryMock.On("Create", mock.Anything).Return(newProduct, nil)
			defer productRepositoryMock.ClearExpectations()

			createdProduct, err := productService.Create(newProduct)
			assert.NoError(test, err)
			assert.Equal(test, newProduct, createdProduct)
		})
	})
}
