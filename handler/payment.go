package handler

import (
	"api/model"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func UploadPayments(c *fiber.Ctx) error {
	fileInfo := new(model.FileInfo)

	// Store the body in the ticket and return an error if encountered
	err := c.BodyParser(fileInfo)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	file, err := c.FormFile("image")
	if err != nil {
		log.Println("Error in uploading Image : ", err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})

	}

	filename := fileInfo.ID
	fileExt := strings.Split(file.Filename, ".")[1]
	image := fmt.Sprintf("%s.%s", filename, fileExt)

	err = c.SaveFile(file, fmt.Sprintf("./images/%s", image))

	if err != nil {
		log.Println("Error in saving Image :", err)
		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})
	}

	return c.JSON(fiber.Map{"status": 201, "message": "Image uploaded successfully as" + image})
}

func GetPayment(c *fiber.Ctx) error {
	filename := c.Query("file")

	if filename == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Filename parameter is required", "data": nil})
	}

	imagePath := fmt.Sprintf("./images/%s", filename)

	// Check if the file exists
	if _, err := os.Stat(imagePath); os.IsNotExist(err) {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "Image not found", "data": nil})
	}

	// Read the image file
	imageData, err := ioutil.ReadFile(imagePath)
	if err != nil {
		log.Println("Error reading image file:", err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Server error", "data": nil})
	}

	// Return the image data as response
	return c.Send(imageData)
}
