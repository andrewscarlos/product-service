package controller

import (
	"context"
	"product-service/internal/product/pb"
)

type ProductGrpcController struct {
	pb.UnimplementedProductControllerServer
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
