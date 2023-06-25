package main

import (
	"github.com/gofiber/fiber/v2"
	accountcontroller "github.com/sudrajad02/ecommerce/controller/accountcontroller"
	"github.com/sudrajad02/ecommerce/database"
)

func init() {
	database.ConncetionDatabase()
}

func main() {
	app := fiber.New()

	api := app.Group("/api")
	book := api.Group("/book")

	book.Get("/", accountcontroller.Index)

	app.Listen(":3000")
}
