package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pllehrman/trivia-example/handlers"
)

func setupRotues(app *fiber.App) {
	app.Get("/", handlers.Home)
	app.Post("/fact", handlers.CreateFact)
}

	