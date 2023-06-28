package authcontroller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sudrajad02/ecommerce/database"
	"github.com/sudrajad02/ecommerce/models"
)

func Login(c *fiber.Ctx) error {
	payload := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	if payload.Username == "" || payload.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Please fill username and password",
		})
	}

	var account models.Account

	if err := database.DB.First(&account, "account_username = ? and account_password = ?", payload.Username, payload.Password).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// claims := jwt.MapClaims{
	// 	"name":  "John Doe",
	// 	"admin": true,
	// 	"exp":   time.Now().Add(time.Hour * 72).Unix(),
	// }

	return c.Status(fiber.StatusOK).JSON(account)

}

func Register(c *fiber.Ctx) error {
	var account models.Account

	if err := c.BodyParser(&account); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := database.DB.Save(&account).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(account)
}

func DetailAccount(c *fiber.Ctx) error {
	payload := struct {
		UserId int `json:"account_id"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	var account models.Account

	database.DB.Where("account_id = ?", payload.UserId).Preload("AccountAddress", "account_address_is_active = ?", 1).Find(&account)

	return c.Status(fiber.StatusOK).JSON(account)
}
