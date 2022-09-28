package main

import (
	"fmt"
	"log"
	"product-service/internal/product/repository/postgres"
	"product-service/internal/product/router"
	"product-service/internal/product/service"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := postgres.NewConnection()
	productRepository := postgres.NewProductPostgresRepository(db)
	productService := service.NewProductService(productRepository)
	app := fiber.New(fiber.Config{})
	routes := app.Group("/v1")
	router.NewRouters(routes, productService)
	fmt.Println("product service")
	app.Listen(":3001")
}
