package api

import (
	"github.com/nao99/spring-boot-3-vs-go/products-go/internal/server"
)

func SetupRoutes(srv *server.Server, productController *ProductController) {
	srv.Get("/api/v1/products/:productId", productController.GetProduct)
	srv.Get("/api/v1/products", productController.GetProducts)

	srv.Post("/api/v1/products", productController.CreateProduct)
}
