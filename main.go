package main

import (
	db "auth_fiber/config"
	routes "auth_fiber/routes"

	"github.com/gofiber/fiber/v2"
)



func main()  {
	db.Connect();
	db.MongoConnection();
	app :=fiber.New(fiber.Config{
		BodyLimit: 10*1024*1024,
	});
	app.Get("/",func(c *fiber.Ctx) error {
		return c.SendString("Hello World!");
	})
	routes.Setup(app)
	app.Listen(":4000");

}

