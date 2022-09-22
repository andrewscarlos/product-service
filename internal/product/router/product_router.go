package router

import (
	"product-service/internal/product/handler"
	"product-service/internal/product/interfaces"

	"github.com/gofiber/fiber/v2"
)

func RegisterRouter(router fiber.Router, productSercice interfaces.ProductServiceInterface) {
	productRouterServiceGroup := router.Group("/product")

	productRouterServiceGroup.Post("/", func(fiberContext *fiber.Ctx) error {
		return handler.CreateProduct(fiberContext, productSercice)
	})

	productRouterServiceGroup.Get("/:id", func(fiberContext *fiber.Ctx) error {
		return handler.GetProductById(fiberContext, productSercice)
	})

	productRouterServiceGroup.Get("/", func(fiberContext *fiber.Ctx) error {
		return handler.GetAllProducts(fiberContext, productSercice)
	})

	productRouterServiceGroup.Put("/:id", func(fiberContext *fiber.Ctx) error {
		return handler.UpdateProducts(fiberContext, productSercice)
	})

	productRouterServiceGroup.Delete("/:id", func(fiberContext *fiber.Ctx) error {
		return handler.DeleteProduct(fiberContext, productSercice)
	})

}
