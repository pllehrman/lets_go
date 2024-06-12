package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pllehrman/trivia-example/database"
)
func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRotues(app)

	app.Listen(":3000")
}

