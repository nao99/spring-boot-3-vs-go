package domain

import (
	"github.com/nao99/spring-boot-3-vs-go/products-go/internal/domain/model"
	"github.com/nao99/spring-boot-3-vs-go/products-go/internal/persistence/api"
)

type ProductService interface {
	GetProducts() *[]model.Product
	GetProductById(id uint64) *model.Product
	Create(category string, name string, description string) *model.Product
}

type ProductServiceImpl struct {
	Repository repository.ProductRepository
}

func (srv *ProductServiceImpl) GetProducts() *[]model.Product {
	return srv.Repository.FindLast20Products()
}

func (srv *ProductServiceImpl) GetProductById(id uint64) *model.Product {
	product := srv.Repository.FindById(id)
	if product == nil {
		return nil
	}

	if product.ID == 0 {
		return nil
	}

	return product
}

func (srv *ProductServiceImpl) Create(category string, name string, description string) *model.Product {
	product := model.Product{Name: name, Category: category, Description: description}
	srv.Repository.Save(&product)

	return &product
}
