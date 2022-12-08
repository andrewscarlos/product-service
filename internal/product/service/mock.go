package service

import (
	"product-service/internal/product"

	"github.com/stretchr/testify/mock"
)

type Mock struct {
	mock.Mock
	ProductServiceInterface
}

func (service *Mock) Create(p *product.Product) (*product.Product, error) {
	args := service.Called(p)
	return args.Get(0).(*product.Product), args.Error(1)
}

func (service *Mock) GetById(ID string) (*product.Product, error) {
	args := service.Called(ID)
	return args.Get(0).(*product.Product), args.Error(1)
}

func (service *Mock) GetAll() ([]*product.Product, error) {
	args := service.Called()
	return args.Get(0).([]*product.Product), args.Error(1)
}

func (service *Mock) Delete(ID string) (bool, error) {
	args := service.Called(ID)
	return args.Get(0).(bool), args.Error(1)
}

func (service *Mock) Uptate(ID string, p *product.Product) (*product.Product, error) {
	args := service.Called(ID, p)
	return args.Get(0).(*product.Product), args.Error(1)
}

func (service *Mock) ClearExpectations() {
	service.ExpectedCalls = nil
}
