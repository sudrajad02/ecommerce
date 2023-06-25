package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sudrajad02/ecommerce/controller/authcontroller"
	"github.com/sudrajad02/ecommerce/controller/productcontroller"
	"github.com/sudrajad02/ecommerce/database"
)

func init() {
	database.ConncetionDatabase()
}

func main() {
	app := fiber.New()

	api := app.Group("/api")
	auth := api.Group("/auth")
	product := api.Group("/product")

	// auth
	auth.Post("/", authcontroller.Login)
	auth.Post("/register/", authcontroller.Register)

	// product
	product.Get("/:category_id?", productcontroller.ListProduct)
	product.Get("/detail/:product_id", productcontroller.DetailProduct)

	app.Listen(":3000")
}
