package controller

import (
	"auth_fiber/config"
	"auth_fiber/models"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)


func LoginUser(c *fiber.Ctx) error {
		cashierId:=c.Params("cashierId");

		var data map[string]string;

		err :=c.BodyParser(&data)

		if(err!=nil){
			return c.Status(400).JSON(fiber.Map{
			"success":false,
			"message":"Invalid Data",
			})
		}


		if data["passcode"]==""{
			return c.Status(400).JSON(fiber.Map{
				"success":false,
				"message":"Passcode Is Required",
				"error":map[string]interface{}{},
			})
		}
		var cashier models.Cashier

		config.DB.Where("id=?",cashierId).First(&cashier)

		if(cashier.Id==0){
			return c.Status(400).JSON(fiber.Map{
				"success":false,
				"message":"Cashier Not Found",
				"error":map[string]interface{}{},
			})
		}
		if cashier.Passcode != data["passcode"]{
				return c.Status(400).JSON(fiber.Map{
				"success":false,
				"message":"Password doesn't matched",
				"error":map[string]interface{}{},
			})
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
			"Issuer":strconv.Itoa(int(cashier.Id)),
			"ExpiresAt":time.Now().Add(time.Hour * 24).Unix(),
		})

		tokenString,err:=token.SignedString([]byte(os.Getenv("JWT_SECRET")))
		if(err!=nil){
			return c.Status(400).JSON(fiber.Map{
				"success":false,
				"message":"Token Create Failed",
			})
		}

		return c.Status(200).JSON(fiber.Map{
				"success":false,
				"message":"Login Success",
				"token":tokenString,
			})		
}
func LogoutUser(c *fiber.Ctx) error {
	return nil;
}
func Passcode(c *fiber.Ctx) error {
	return nil;
}
