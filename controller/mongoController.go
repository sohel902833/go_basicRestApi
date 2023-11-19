package controller

import "github.com/gofiber/fiber/v2"


func CreateMovie(c *fiber.Ctx) error{
	return c.JSON(fiber.Map{
		 "mongo":true,
	})
}