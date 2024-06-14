package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/lib/pq"
	"github.com/shivampri2002/Coding_Judge_Platform/database"
	"github.com/shivampri2002/Coding_Judge_Platform/router"
)

func main() {
	// Connect to database
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	// call the New() method - used to instantiate a new Fiber App
	app := fiber.New()

	// Middleware
	app.Use(logger.New())

	router.SetupRoutes(app)

	// Last middleware to match anything
	app.Use(func(c *fiber.Ctx) error {
		c.SendStatus(404) // => 404 "Not Found"
		return c.SendString("404 Not Found")
	})

	// listen on port 3000
	app.Listen(":3000")
}
