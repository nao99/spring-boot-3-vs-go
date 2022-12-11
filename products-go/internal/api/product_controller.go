package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nao99/spring-boot-3-vs-go/products-go/internal/api/model"
	"github.com/nao99/spring-boot-3-vs-go/products-go/internal/domain"
	"strconv"
)

type ProductController struct {
	Service domain.ProductService
}

func (ctrl *ProductController) GetProduct(ctx *fiber.Ctx) error {
	productIdString := ctx.Params("productId")

	productId, err := strconv.ParseUint(productIdString, 10, 32)
	if err != nil {
		return err
	}

	product := ctrl.Service.GetProductById(productId)
	if product == nil {
		return fiber.NewError(fiber.StatusNotFound, "Product not found")
	}

	return ctx.JSON(product)
}

func (ctrl *ProductController) GetProducts(ctx *fiber.Ctx) error {
	products := ctrl.Service.GetProducts()
	return ctx.JSON(products)
}

func (ctrl *ProductController) CreateProduct(ctx *fiber.Ctx) error {
	createDto := &dto.CreateProductDto{}

	err := ctx.BodyParser(&createDto)
	if err != nil {
		return err
	}

	product := ctrl.Service.Create(createDto.Category, createDto.Name, createDto.Description)

	return ctx.JSON(product)
}
