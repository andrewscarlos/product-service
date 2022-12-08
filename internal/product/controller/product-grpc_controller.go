package controller

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"product-service/internal/product"
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

func (pc *ProductGrpcController) Create(ctx context.Context, p *pb.Product) (*pb.Product, error) {
	newProduct := new(product.Product)
	newProduct.ID = uuid.NewV4()
	newProduct.Name = p.Name
	newProduct.Type = p.Type
	newProduct.Value = float64(p.Value)
	productCreated, err := pc.service.Create(newProduct)
	if err != nil {
		return nil, err
	}
	product := &pb.Product{
		Value: float32(productCreated.Value),
		Name:  productCreated.Name,
		Type:  productCreated.Type,
	}
	return product, nil
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
