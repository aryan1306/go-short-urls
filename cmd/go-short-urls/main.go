package main

import (
	"context"
	"log"
	"os"

	"github.com/aryan1306/go-short-urls/internal/redisClient"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var ctx = context.Background()

func main() {
	err := godotenv.Load("../../.env")
	redisUrl := os.Getenv("REDIS_URL")
  if err != nil {
    log.Fatal("Error loading .env file", err)
  }
	app := fiber.New()
	_ = redisClient.Init(redisUrl)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":8000")
}