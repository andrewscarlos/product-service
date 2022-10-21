package mongodb

import (
	"log"
	"product-service/internal/product"
	"product-service/internal/product/repository"

	"gopkg.in/mgo.v2"
)

func NewConnection() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		log.Fatal("Can not connect with mongodb")
	}

	return session
}

// *mgo.Database
type ProductMongoDbRepository struct {
	mongo *mgo.Collection
}

func NewProductMongoDbRepository(db *mgo.Session) repository.ProductRepositoryInterface {
	return &ProductMongoDbRepository{
		mongo: db.DB("").C("products"),
	}
}

func (repository *ProductMongoDbRepository) Create(product *product.Product) (*product.Product, error) {
	repository.mongo.Insert(product)
	return nil, nil
}

func (repository *ProductMongoDbRepository) GetById(ID string) (*product.Product, error) {
	return nil, nil
}

func (repository *ProductMongoDbRepository) GetAll() ([]*product.Product, error) {
	return nil, nil
}

func (repository *ProductMongoDbRepository) Delete(ID string) (bool, error) {
	return false, nil
}

func (repository *ProductMongoDbRepository) Uptate(ID string, product *product.Product) (*product.Product, error) {
	return nil, nil
}
