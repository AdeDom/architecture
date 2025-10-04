package adapters

import (
	"architecture/core/product"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

// Primary adapter
type HttpProductHandler struct {
	service core.ProductService
}

func NewHttpProductHandler(service core.ProductService) *HttpProductHandler {
	return &HttpProductHandler{service: service}
}

func (h *HttpProductHandler) CreateProduct(c *fiber.Ctx) error {
	var product core.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "Invalid request",
		})
	}

	result, err := h.service.CreateProduct(product)
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

func (h *HttpProductHandler) GetProductById(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  "Invalid request",
		})
	}

	product, err := h.service.GetProductById(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(product)
}

func (h *HttpProductHandler) GetProducts(c *fiber.Ctx) error {
	name := c.Query("name")
	price, _ := strconv.ParseFloat(c.Query("price"), 64)

	products, err := h.service.GetProducts(name, price)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	return c.JSON(products)
}
