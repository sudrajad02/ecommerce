package productcontroller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sudrajad02/ecommerce/database"
	"github.com/sudrajad02/ecommerce/models"
)

func ListProduct(c *fiber.Ctx) error {
	var product []models.Product

	queryValue := c.Params("category_id")

	if queryValue == "" {
		database.DB.Preload("ProductCategory").Find(&product)
	} else {
		database.DB.Preload("ProductCategory").Where("product_category_id = ?", queryValue).Find(&product)
	}

	return c.Status(fiber.StatusOK).JSON(product)
}

func DetailProduct(c *fiber.Ctx) error {
	var product models.Product

	queryValueProductId := c.Params("product_id")

	database.DB.Where("product_id = ?", queryValueProductId).Find(&product)

	return c.Status(fiber.StatusOK).JSON(product)
}
