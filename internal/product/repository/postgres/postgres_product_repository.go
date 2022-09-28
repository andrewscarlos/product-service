package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"product-service/internal/product"
	"product-service/internal/product/repository"
)

func NewConnection() *gorm.DB {
	host := os.Getenv("PRODUCT_POSTGRES_HOST")
	port := os.Getenv("PRODUCT_POSTGRES_PORT")
	postgres_db := os.Getenv("PRODUCT_POSTGRES_DB")
	postgres_user := os.Getenv("PRODUCT_POSTGRES_USER")
	postgres_password := os.Getenv("PRODUCT_POSTGRES_PASSWORD")

	stringConnection := "host=" + host + " user=" + postgres_user + " password=" + postgres_password + " dbname=" + postgres_db + " port=" + port + " sslmode=disable" + " TimeZone=America/Sao_Paulo"

	gorm, err := gorm.Open(postgres.Open(stringConnection), &gorm.Config{})
	if err != nil {
		log.Fatal("Can not connect with postgres")
	}
	gorm.AutoMigrate(product.Product{})
	return gorm
}

type ProductPostgresRepository struct {
	postgres *gorm.DB
}

func NewProductPostgresRepository(db *gorm.DB) repository.ProductRepositoryInterface {
	return &ProductPostgresRepository{
		postgres: db,
	}
}

func (repository *ProductPostgresRepository) Create(productData *product.Product) (*product.Product, error) {
	productCreated := repository.postgres.Create(productData)

	if productCreated.Error != nil {
		return nil, productCreated.Error
	}

	return productData, nil
}

func (repository *ProductPostgresRepository) GetById(ID string) (*product.Product, error) {

	product := &product.Product{}

	productFound := repository.postgres.First(product, "id = ?", ID)

	if productFound.Error != nil {
		return nil, productFound.Error
	}

	return product, nil
}

func (repository *ProductPostgresRepository) GetAll() ([]*product.Product, error) {
	var products []*product.Product

	err := repository.postgres.Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (repository *ProductPostgresRepository) Delete(ID string) (bool, error) {
	deletedProduct := repository.postgres.Delete(ID)

	if deletedProduct.Error != nil {
		return false, deletedProduct.Error
	}
	return true, nil
}

func (repository *ProductPostgresRepository) Uptate(ID string, product *product.Product) (*product.Product, error) {
	productUpdated := repository.postgres.Update(ID, product)

	if productUpdated.Error != nil {
		return nil, productUpdated.Error
	}

	return product, nil
}
