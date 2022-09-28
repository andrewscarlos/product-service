package router

import (
	"product-service/internal/product/controller"
	"product-service/internal/product/service"

	"github.com/gofiber/fiber/v2"
)

func NewRouters(router fiber.Router, productSercice service.ProductServiceInterface) {
	productRouterServiceGroup := router.Group("/product")

	productRouterServiceGroup.Post("/", func(fiberContext *fiber.Ctx) error {
		return controller.CreateProduct(fiberContext, productSercice)
	})

	productRouterServiceGroup.Get("/:id", func(fiberContext *fiber.Ctx) error {
		return controller.GetProductById(fiberContext, productSercice)
	})

	productRouterServiceGroup.Get("/", func(fiberContext *fiber.Ctx) error {
		return controller.GetAllProducts(fiberContext, productSercice)
	})

	productRouterServiceGroup.Put("/:id", func(fiberContext *fiber.Ctx) error {
		return controller.UpdateProducts(fiberContext, productSercice)
	})

	productRouterServiceGroup.Delete("/:id", func(fiberContext *fiber.Ctx) error {
		return controller.DeleteProduct(fiberContext, productSercice)
	})

}
