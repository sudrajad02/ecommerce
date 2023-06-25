package bookcontroller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sudrajad02/ecommerce/database"
	"github.com/sudrajad02/ecommerce/models"
)

func Index(c *fiber.Ctx) error {
	var account []models.Account

	database.DB.Find(&account)

	return c.Status(fiber.StatusOK).JSON(account)
}
