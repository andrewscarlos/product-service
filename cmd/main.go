package main

import (
	"fmt"
	"log"
	"net"
	//"os"
	"product-service/internal/product/pb"
	"product-service/internal/product/repository"

	postgresImplementation "product-service/pkg/repository/postgres"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

	lis, err := net.Listen("tcp", "localhost:5000")

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterProductControllerServer(grpcServer, controller.NewProductGrpcController(productService))

	reflection.Register(grpcServer)

	go func() {
		fmt.Println("product service running on grpc port 5000")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Could not serve: %v", err)
		}

	}()
	fmt.Println("product service running on rest port 3001")
	app.Listen(":3001")
}
