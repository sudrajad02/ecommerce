package cartcontroller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sudrajad02/ecommerce/database"
	"github.com/sudrajad02/ecommerce/models"
)

func ListCart(c *fiber.Ctx) error {
	payload := struct {
		UserId int `json:"account_id"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	var cart []models.Cart

	database.DB.Where("cart_account_id = ?", payload.UserId).Preload("Product.ProductCategory").Preload("Account").Find(&cart)

	return c.Status(fiber.StatusOK).JSON(cart)
}

func AddCart(c *fiber.Ctx) error {
	var cart models.Cart

	if err := c.BodyParser(&cart); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := database.DB.Save(&cart).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(cart)
}
