package controller

import (
	"auth_fiber/config"
	"auth_fiber/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateCashier(c *fiber.Ctx) error {
	var data map[string]string;

	 err:=c.BodyParser(&data);
	if(err!=nil){
		return c.Status(400).JSON(
			fiber.Map{
				"success":false,
				"message":"Invalid Information",
			},
		)
	}
	 errors := make(map[string]string)
	dataExists :=len(data)>0;
	if(!dataExists){
		errors["message"]="No Data Found";
	}
	if dataExists && data["name"]==""{
		errors["name"]="Name is required";
	}

	if dataExists &&  data["passcode"]==""{
		errors["passcode"]="Pass code is required"
	}

	if(len(errors)!=0){
		return c.Status(400).JSON(
			fiber.Map{
				"success":false,
				"message":"Invalid Information",
				"errors":errors,
			},
		)
	}

	cashier:=models.Cashier{
		Name: data["name"],
		Passcode: data["passcode"],
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	config.DB.Create(&cashier);

	return c.Status(200).JSON(fiber.Map{
		"success":true,
		"message":"Cashier Created",
		"cashier":cashier,
	})


}
func GetCashierList(c *fiber.Ctx) error {
	var cashier []models.Cashier
	limit,err:=strconv.Atoi(c.Query("limit"));
	if(err!=nil || limit<=0){
		limit=5;
	}
	skip,_:=strconv.Atoi(c.Query("skip"));
	if(err!=nil || skip<=0){
		skip=0;
	}

	var count int64;

	config.DB.Select("*").Limit(limit).Offset(skip).Find(&cashier).Count(&count)

	return c.Status(200).JSON(fiber.Map{
		"success":true,
		"message":"Cashier list api",
		"data":cashier,
		"total_data":count,
	})
}
func GetSingleCashierInfo(c *fiber.Ctx) error {
	cashierId:=c.Params("cashierId");

	var cashier models.Cashier;

	config.DB.Select("*").Where("id=?",cashierId).First(&cashier)

	if(cashier.Id==0){
		return c.Status(200).JSON(fiber.Map{
			"success":false,
			"message":"No Cashier Found",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success":true,
		"message":"Cashier Info",
		"data":cashier,
	})
}
func EditCashier(c *fiber.Ctx) error {
	cashierId:=c.Params("cashierId");

	var cashier models.Cashier

	config.DB.Find(&cashier,"id=?",cashierId)

	if(cashier.Name==""){
		return c.Status(200).JSON(fiber.Map{
		"success":false,
		"message":"Cashier Not Found",
	})
	}

	var updateCashier models.Cashier
	err :=c.BodyParser(&updateCashier)
	if(err!=nil){
		return err;
	}
	if(updateCashier.Name==""){
		return c.Status(404).JSON(fiber.Map{
			"success":false,
			"message":"Nothing found to update",
		})
	}

	cashier.Name=updateCashier.Name

	config.DB.Save(&cashier)


	return c.Status(200).JSON(fiber.Map{
		"success":true,
		"message":"Cashier Updated",
		"data":cashier,
	})
}
func DeleteCashier(c *fiber.Ctx) error {
	cashierId:=c.Params("cashierId");

	var cashier models.Cashier

	config.DB.Where("id=?",cashierId).First(&cashier)

	if(cashier.Id==0){
		return c.Status(200).JSON(fiber.Map{
			"success":false,
			"message":"Cashier Not Found",
		})
  	}

 config.DB.Where("id=?",cashierId).Delete(&cashier);


	return c.Status(200).JSON(fiber.Map{
			"success":false,
			"message":"Cashier Deleted",
			
		})

}