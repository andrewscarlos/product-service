package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"product-service/internal/product/repository"
	"product-service/internal/product/repository/mongodb"
	//"product-service/internal/product/repository/postgres"
	"product-service/internal/product/router"
	"product-service/internal/product/service"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//postgresdb := postgres.NewConnection()
	//postgresRepository := postgres.NewProductPostgresRepository(postgresdb)

	mongodbCon := mongodb.NewConnection()
	mongodbRepository := mongodb.NewProductMongoDbRepository(mongodbCon)

	productRepository := repository.NewProductRepository(mongodbRepository)
	productService := service.NewProductService(productRepository)
	app := fiber.New(fiber.Config{})
	routes := app.Group("/v1")
	router.NewRouters(routes, productService)
	fmt.Println("product service")
	app.Listen(":3001")
}
