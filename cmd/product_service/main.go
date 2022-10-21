package main

import (
	"fmt"
	"log"
	"product-service/internal/product/repository"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	postgresImplementation"product-service/pkg/repository/postgres"

	//mongoDbImplementation"product-service/pkg/repository/mongodb"
	"product-service/internal/product/controller"

	"product-service/internal/product/service"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	postgresdb := postgresImplementation.NewConnection()
	postgresRepository := postgresImplementation.NewProductPostgresRepository(postgresdb)

	// mongodbCon := mongodb.NewConnection()
	// mongodbRepository := mongodb.NewProductMongoDbRepository(mongodbCon)

	productRepository := repository.NewProductRepository(postgresRepository)
	productService := service.NewProductService(productRepository)
	app := fiber.New(fiber.Config{})
	routes := app.Group("/v1")
	controller.NewRouters(routes, productService)
	fmt.Println("product service")
	app.Listen(":3001")
}
