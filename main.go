package main

import (
	"github.com/gofiber/fiber/v2"
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

	auth.Post("/", authcontroller.Login)
	auth.Post("/register/", authcontroller.Register)

	app.Listen(":3000")
}
