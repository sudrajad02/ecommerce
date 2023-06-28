package checkoutcontroller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sudrajad02/ecommerce/database"
	"github.com/sudrajad02/ecommerce/models"
)

func ListCheckout(c *fiber.Ctx) error {
	payload := struct {
		UserId int `json:"account_id"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	var checkout []models.Checkout

	database.DB.Where("checkout_account_id = ?", payload.UserId).Preload("Product.ProductCategory").Preload("Account").Preload("AccountAddress").Preload("Product").Find(&checkout)

	return c.Status(fiber.StatusOK).JSON(checkout)
}

func AddCheckout(c *fiber.Ctx) error {
	var checkout models.Checkout
	var cart models.Cart

	if err := c.BodyParser(&checkout); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := database.DB.Save(&checkout).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	payload := struct {
		CartId int `json:"cart_id"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := database.DB.Where("cart_id = ?", payload.CartId).Delete(&cart).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(checkout)
}
