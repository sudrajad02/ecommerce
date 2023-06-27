package cartcontroller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sudrajad02/ecommerce/database"
	"github.com/sudrajad02/ecommerce/models"
	"gorm.io/gorm/clause"
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
	payload := struct {
		ProductId int `json:"product_id"`
		AccountId int `json:"account_id"`
		Ammount   int `json:"amount"`
		Price     int `json:"price"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var cart models.Cart

	database.DB.Where("cart_product_id = ? AND cart_account_id = ?", payload.ProductId, payload.AccountId).Preload("Product.ProductCategory").Preload("Account").Find(&cart)

	if cart.CartProductId == 0 {
		new_body := models.Cart{
			CartProductId: payload.ProductId,
			CartAccountId: payload.AccountId,
			AmountProduct: cart.AmountProduct + payload.Ammount,
			TotalPrice:    cart.TotalPrice + (payload.Price * payload.Ammount),
		}

		if err := database.DB.Save(&new_body).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
	} else {
		new_body := models.Cart{
			CartProductId: payload.ProductId,
			CartAccountId: payload.AccountId,
			AmountProduct: cart.AmountProduct + payload.Ammount,
			TotalPrice:    cart.TotalPrice + (payload.Price * payload.Ammount),
		}

		if err := database.DB.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "CartProductId"}, {Name: "CartAccountId"}},
			DoUpdates: clause.AssignmentColumns([]string{"amount_product", "total_price"}),
		}).Save(&new_body).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(cart)
}

func DeleteCart(c *fiber.Ctx) error {
	var cart models.Cart

	if err := database.DB.Where("cart_id = ?", c.Params("id")).Delete(&cart).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(cart)
}
