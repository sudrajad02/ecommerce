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
	auth.Post("/login", authcontroller.Login)
	auth.Post("/register", authcontroller.Register)
	auth.Post("/detail-user", authcontroller.SessionChecker, authcontroller.DetailAccount)

	// product
	product.Get("/:category_id?", productcontroller.ListProduct)
	product.Get("/detail/:product_id", productcontroller.DetailProduct)

	// cart
	cart.Post("/", authcontroller.SessionChecker, cartcontroller.ListCart)
	cart.Post("/add", authcontroller.SessionChecker, cartcontroller.AddCart)
	cart.Delete("/delete/:id", authcontroller.SessionChecker, cartcontroller.DeleteCart)

	// checkout
	checkout.Post("/", authcontroller.SessionChecker, checkoutcontroller.ListCheckout)
	checkout.Post("/add", authcontroller.SessionChecker, checkoutcontroller.AddCheckout)

	// payment
	payment.Post("/", authcontroller.SessionChecker, deliverycontroller.UpdateStatusPayment)

	// delivery
	delivery.Post("/", authcontroller.SessionChecker, deliverycontroller.UpdateStatusDelivery)

	app.Listen(":3000")
}
