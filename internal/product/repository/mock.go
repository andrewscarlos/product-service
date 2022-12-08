package repository

import (
	"product-service/internal/product"

	"github.com/stretchr/testify/mock"
)

type Mock struct {
	mock.Mock
	ProductRepositoryInterface
}

func (repository *Mock) Create(p *product.Product) (*product.Product, error) {
	args := repository.Called(p)
	return args.Get(0).(*product.Product), args.Error(1)
}

func (repository *Mock) GetById(ID string) (*product.Product, error) {
	args := repository.Called(ID)
	return args.Get(0).(*product.Product), args.Error(1)
}

func (repository *Mock) GetAll() ([]*product.Product, error) {
	args := repository.Called()
	return args.Get(0).([]*product.Product), args.Error(1)
}

func (repository *Mock) Delete(ID string) (bool, error) {
	args := repository.Called(ID)
	return args.Get(0).(bool), args.Error(1)
}

func (repository *Mock) Uptate(ID string, p *product.Product) (*product.Product, error) {
	args := repository.Called(ID, p)
	return args.Get(0).(*product.Product), args.Error(1)
}

func (repository *Mock) ClearExpectations() {
	repository.ExpectedCalls = nil
}
