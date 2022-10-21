package http

import (
	"encoding/json"
	"product-service/internal/product"
	"product-service/internal/product/service"

	"github.com/gofiber/fiber/v2"
	uuid "github.com/satori/go.uuid"
)

func CreateProduct(fiberContext *fiber.Ctx, productService service.ProductServiceInterface) error {
	newProduct := new(product.Product)
	newProduct.ID = uuid.NewV4()
	if err := json.Unmarshal(fiberContext.Body(), newProduct); err != nil {
		return fiberContext.Status(fiber.StatusBadRequest).JSON(map[string]string{
			"message": err.Error(),
		})
	}
	productCreated, err := productService.Create(newProduct)
	if err != nil {
		return fiberContext.Status(fiber.StatusBadRequest).JSON(map[string]string{
			"message": err.Error(),
		})
	}
	fiberContext.Status(fiber.StatusCreated).JSON(productCreated)
	return nil
}

func GetProductById(fiberContext *fiber.Ctx, productService service.ProductServiceInterface) error {
	return nil
}

func GetAllProducts(fiberContext *fiber.Ctx, productService service.ProductServiceInterface) error {
	return nil
}

func UpdateProducts(fiberContext *fiber.Ctx, productService service.ProductServiceInterface) error {
	return nil
}

func DeleteProduct(fiberContext *fiber.Ctx, productService service.ProductServiceInterface) error {
	return nil
}
