package routes

import (
	controller "auth_fiber/controller"

	"github.com/gofiber/fiber/v2"
)


func MongoRoutes(app *fiber.App){

	mongoRouteGroups:=app.Group("/mongo")
	mongoRouteGroups.Post("/movie",controller.CreateMovie)
}