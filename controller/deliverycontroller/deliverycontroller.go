package deliverycontroller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sudrajad02/ecommerce/database"
	"github.com/sudrajad02/ecommerce/models"
)

func UpdateStatusPayment(c *fiber.Ctx) error {
	// can use web hook for payment gateway
	var checkout models.Checkout

	payload := struct {
		CheckoutId int `json:"checkout_id"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	if err := database.DB.Model(&checkout).Where("checkout_id = ?", payload.CheckoutId).Update("checkout_status_delivery", "Sudah Dibayar, Menunggu Dikemas!").Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := database.DB.Preload("Account").Preload("AccountAddress").Preload("Product").Where("checkout_id = ?", payload.CheckoutId).Find(&checkout).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(checkout)
}

func UpdateStatusDelivery(c *fiber.Ctx) error {
	var checkout models.Checkout

	payload := struct {
		CheckoutId int `json:"checkout_id"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	if err := database.DB.Preload("Account").Preload("AccountAddress").Preload("Product").Where("checkout_id = ?", payload.CheckoutId).Find(&checkout).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if checkout.CheckoutStatusDelivery == "Sudah Dibayar, Menunggu Dikemas!" {
		if err := database.DB.Model(&checkout).Where("checkout_id = ?", payload.CheckoutId).Update("checkout_status_delivery", "Pesanan Sedang Diantar!").Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
	} else if checkout.CheckoutStatusDelivery == "Pesanan Sedang Diantar!" {
		if err := database.DB.Model(&checkout).Where("checkout_id = ?", payload.CheckoutId).Update("checkout_status_delivery", "Pesanan Selesai!").Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
	}

	if err := database.DB.Preload("Account").Preload("AccountAddress").Preload("Product").Where("checkout_id = ?", payload.CheckoutId).Find(&checkout).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(checkout)
}
