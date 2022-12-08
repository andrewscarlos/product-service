package controller

import (
	"product-service/internal/product/service"

	"github.com/gofiber/fiber/v2"
)

func NewRouters(router fiber.Router, productSercice service.ProductServiceInterface) {
	productRouterServiceGroup := router.Group("/product")
	// HealthCheck godoc
	// @Summary Show the status of server.
	// @Description create new product.
	// @Tags root
	// @Accept */*
	// @Produce json
	// @Success 200 {object} map[string]interface{}
	// @Router / [post]
	productRouterServiceGroup.Post("/", func(fiberContext *fiber.Ctx) error {
		return CreateProduct(fiberContext, productSercice)
	})

	productRouterServiceGroup.Get("/:id", func(fiberContext *fiber.Ctx) error {
		return GetProductById(fiberContext, productSercice)
	})

	productRouterServiceGroup.Get("/", func(fiberContext *fiber.Ctx) error {
		return GetAllProducts(fiberContext, productSercice)
	})

	productRouterServiceGroup.Put("/:id", func(fiberContext *fiber.Ctx) error {
		return UpdateProducts(fiberContext, productSercice)
	})

	productRouterServiceGroup.Delete("/:id", func(fiberContext *fiber.Ctx) error {
		return DeleteProduct(fiberContext, productSercice)
	})

}
