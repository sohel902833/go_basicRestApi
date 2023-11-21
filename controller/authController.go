package controller

import (
	"auth_fiber/config"
	"auth_fiber/models"
	"fmt"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)


func LoginUser(c *fiber.Ctx) error {
		return nil;	
}

func SignupUser(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Failed to parse image",
		})
	}
	// Read the image file
	fileContent, err := file.Open()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to read image file",
		})
	}
	defer fileContent.Close()
	timestampMillis :=strconv.Itoa(int(time.Now().UnixMilli()));

	// Create a unique filename based on the user's name or any other identifier
	filename := fmt.Sprintf("%s_%s_%s", c.FormValue("name"),timestampMillis, file.Filename)

	// Save the image to the file system
	savePath := filepath.Join("uploads", filename)
	if err := c.SaveFile(file, savePath); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": fmt.Sprintf("Failed to save image to the file system: %v", err),
		})
	}
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to save image to the file system",
			"err":err,
		})
	}
	newUser :=models.UserModel{
		Email: c.FormValue("email"),
		Image: savePath,
		SignupAt: time.Now(),
	}

	// Save the user to MongoDB
	savedUser,dbError := config.USER.InsertOne(nil, newUser)
	if(dbError!=nil){
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to save value into database",
			"err":err,
		})
	}
	// Your logic for saving additional user information to MongoDB, if needed

	return c.JSON(fiber.Map{
		"success": true,
		"message": "User registered successfully",
		"user": savedUser,
	})
}
