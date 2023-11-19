package routes

import (
	controller "auth_fiber/controller"

	"github.com/gofiber/fiber/v2"
)


func Setup(app *fiber.App){

  
	app.Post("/cashiers/:cashierId/login",controller.LoginUser)

	app.Get("/cashiers/:cashierId/logout",controller.LogoutUser)

	app.Post("/cashiers/:cashierId/passcode",controller.Passcode)


	app.Post("/cashiers",controller.CreateCashier)
	app.Get("/cashiers",controller.GetCashierList)
	app.Get("/cashiers/:cashierId",controller.GetSingleCashierInfo)
	app.Delete("/cashiers/:cashierId",controller.DeleteCashier)

	app.Put("/cashiers/:cashierId",controller.EditCashier)
    MongoRoutes(app);
}