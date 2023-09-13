package main

import (
	"os"

	"github.com/ayowilfred95/api"
	"github.com/ayowilfred95/database"
	"github.com/ayowilfred95/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	err := run()

	if err != nil {
		panic(err)
	}
}

func run() error {
	// initiate env
	err := api.LoadEnv()
	if err != nil {
		return err
	}

	// init db
	err = database.InitDB()
	if err != nil {
		return err
	}

	// defer closing db
	defer database.CloseDB()


	// create an app
	// Fiber framework is used just like express in nodejs
	app := fiber.New()

	// add basic middleware
	app.Use(logger.New())
	app.Use(recover.New())

	// add routes
	router.SetupRoutes(app)


	// start server
	var port string

	if port = os.Getenv("PORT"); port == "" {
		port = "8080"
	}
	app.Listen(":" +port)

	return nil

}
