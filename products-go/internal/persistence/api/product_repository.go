package repository

import (
	"github.com/nao99/spring-boot-3-vs-go/products-go/internal/domain/model"
	"gorm.io/gorm"
)

type ProductRepository interface {
	FindLast20Products() *[]model.Product
	FindById(id uint64) *model.Product
	Save(product *model.Product)
}

type ProductRepositoryImpl struct {
	Session *gorm.DB
}

func (repo *ProductRepositoryImpl) FindLast20Products() *[]model.Product {
	var products *[]model.Product

	repo.Session.Order("id desc").Limit(20).Find(&products)

	return products
}

func (repo *ProductRepositoryImpl) FindById(id uint64) *model.Product {
	product := model.Product{}
	repo.Session.Find(&product, id)

	return &product
}

func (repo *ProductRepositoryImpl) Save(product *model.Product) {
	repo.Session.Save(&product)
}
