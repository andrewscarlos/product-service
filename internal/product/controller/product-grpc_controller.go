package controller

import (
	"context"
	"product-service/internal/product/pb"
	"product-service/internal/product/service"
)

type ProductGrpcController struct {
	pb.UnimplementedProductControllerServer
	service service.ProductServiceInterface
}

type ProductGrpcControllerInterface interface {
	Create(context.Context, *pb.Product) (*pb.Product, error)
	GetById(context.Context, *pb.ID) (*pb.Product, error)
	GetAll(context.Context, *pb.Void) (*pb.Products, error)
	Delete(context.Context, *pb.ID) (*pb.Product, error)
	Uptate(context.Context, *pb.Product) (*pb.Product, error)
}

func NewProductGrpcController(productService service.ProductServiceInterface) *ProductGrpcController {
	return &ProductGrpcController{
		service: productService,
	}
}

func (*ProductGrpcController) Create(context.Context, *pb.Product) (*pb.Product, error) {
	return nil, nil
}

func GetById(context.Context, *pb.ID) (*pb.Product, error) {
	return nil, nil
}

func (*ProductGrpcController) GetAll(context.Context, *pb.Void) (*pb.Products, error) {
	return nil, nil
}

func (*ProductGrpcController) Delete(context.Context, *pb.ID) (*pb.Product, error) {
	return nil, nil
}

func (*ProductGrpcController) Uptate(context.Context, *pb.Product) (*pb.Product, error) {
	return nil, nil
}
