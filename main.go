package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sudrajad02/ecommerce/controller/authcontroller"
	"github.com/sudrajad02/ecommerce/controller/cartcontroller"
	"github.com/sudrajad02/ecommerce/controller/checkoutcontroller"
	"github.com/sudrajad02/ecommerce/controller/deliverycontroller"
	"github.com/sudrajad02/ecommerce/controller/productcontroller"
	"github.com/sudrajad02/ecommerce/database"
)

func init() {
	database.ConncetionDatabase()
}

func main() {
	app := fiber.New()

	// grouping endpoint
	api := app.Group("/api")
	auth := api.Group("/auth")
	product := api.Group("/product")
	cart := api.Group("/cart")
	checkout := api.Group("/checkout")
	payment := api.Group("/payment")
	delivery := api.Group("/delivery")

	// auth
	auth.Post("/", authcontroller.Login)
	auth.Post("/register", authcontroller.Register)
	auth.Post("/detail-user", authcontroller.DetailAccount)

	// product
	product.Get("/:category_id?", productcontroller.ListProduct)
	product.Get("/detail/:product_id", productcontroller.DetailProduct)

	// cart
	cart.Post("/", cartcontroller.ListCart)
	cart.Post("/add", cartcontroller.AddCart)
	cart.Delete("/delete/:id", cartcontroller.DeleteCart)

	// checkout
	checkout.Post("/", checkoutcontroller.AddCheckout)

	// payment
	payment.Post("/", deliverycontroller.UpdateStatusPayment)

	// delivery
	delivery.Post("/", deliverycontroller.UpdateStatusDelivery)

	app.Listen(":3000")
}
