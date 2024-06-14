package handler

import (
	"log"
	// "database/sql"
	"github.com/gofiber/fiber/v2"
	// "github.com/shivampri2002/Coding_Judge_Platform/database"

)

func NewApiHandler(c *fiber.Ctx) error {
	log.Println(c.Body())
	c.Status(200)

	return c.SendString("No error")
}