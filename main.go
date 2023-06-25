package main

import (
	"github.com/gofiber/fiber/v2"
	accountcontroller "github.com/sudrajad02/ecommerce/controller/accountcontroller"
	"github.com/sudrajad02/ecommerce/controller/authcontroller"
	"github.com/sudrajad02/ecommerce/database"
)

func init() {
	database.ConncetionDatabase()
}

func main() {
	app := fiber.New()

	api := app.Group("/api")
	auth := api.Group("/auth")
	book := api.Group("/book")

	auth.Post("/", authcontroller.Login)
	auth.Post("/register/", authcontroller.Register)
	book.Get("/", accountcontroller.Index)

	app.Listen(":3000")
}
