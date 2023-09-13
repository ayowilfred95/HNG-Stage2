package router

import (
	"github.com/ayowilfred95/controller"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	persons := app.Group("/api")

	persons.Post("/", controller.CreatePerson)
	persons.Get("/:userId", controller.GetPerson)
	persons.Get("/", controller.GetPersons)
	persons.Put("/:userId", controller.UpdatePerson)
	persons.Delete("/:userId", controller.DeletePerson)
}
