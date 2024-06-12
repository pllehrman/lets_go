package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/pllehrman/trivia-example/database"
	"github.com/pllehrman/trivia-example/models"
)

func Home(c *fiber.Ctx) error {
	log.Println("Hit Home!")
	return c.SendString("Hello World")
}

func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	err := c.BodyParser(fact)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	log.Println("Hit CreateFact!")
	database.DB.Db.Create(&fact)

	return c.Status(200).JSON(fact)
}