package adapters

import (
	"architecture/data/entities"
	"architecture/domain/usecases"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type ProductController struct {
	createProductUseCase  usecases.CreateProductUseCase
	getProductByIdUseCase usecases.GetProductByIdUseCase
	getProductsUseCase    usecases.GetProductsUseCase
}

func NewProductController(
	createProductUseCase usecases.CreateProductUseCase,
	getProductByIdUseCase usecases.GetProductByIdUseCase,
	getProductsUseCase usecases.GetProductsUseCase,
) *ProductController {
	return &ProductController{
		createProductUseCase:  createProductUseCase,
		getProductByIdUseCase: getProductByIdUseCase,
		getProductsUseCase:    getProductsUseCase,
	}
}

func (h *ProductController) CreateProduct(c *fiber.Ctx) error {
	var product entities.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "Invalid request",
		})
	}

	result, err := h.createProductUseCase.Execute(product)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   &result,
	})
}

func (uc *ProductController) GetProductById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "Invalid request",
		})
	}

	product, err := uc.getProductByIdUseCase.Execute(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(product)
}

func (h *ProductController) GetProducts(c *fiber.Ctx) error {
	name := c.Query("name")
	price, _ := strconv.ParseFloat(c.Query("price"), 64)

	products, err := h.getProductsUseCase.Execute(name, price)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(products)
}
