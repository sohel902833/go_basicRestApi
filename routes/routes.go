package routes

import (
	controller "auth_fiber/controller"

	"github.com/gofiber/fiber/v2"
)


func Setup(app *fiber.App){

  
	app.Post("/login",controller.LoginUser)
	app.Post("/signup",controller.SignupUser)
    // MongoRoutes(app);
}