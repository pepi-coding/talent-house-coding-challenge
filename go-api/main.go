package main

import (
	"log"
	"os"

	"github.com/pepi-coding/talent-house-coding-challenge/matrix"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main(){
	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default env vars.")
	}

	apiVersion := os.Getenv("API_VERSION")
	if apiVersion == "" {
		apiVersion = "v1" // default fallback
	}

	app := fiber.New()

	api := app.Group("/api/" + apiVersion)

	api.Post("/qr", matrix.QRHandler)

	app.Listen(":" + os.Getenv("PORT"))
}
