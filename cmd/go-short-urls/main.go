package main

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/aryan1306/go-short-urls/internal/randomString"
	"github.com/aryan1306/go-short-urls/internal/redisClient"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var ctx = context.Background()

type PostUrl struct {
	Url string `json:"url"`
}

func main() {
	err := godotenv.Load("../../.env")
	redisUrl := os.Getenv("REDIS_URL")
  if err != nil {
    log.Fatal("Error loading .env file", err)
  }
	app := fiber.New()
	client := redisClient.Init(redisUrl)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/shorten", func(c *fiber.Ctx) error {
		p := new(PostUrl)
		if err := c.BodyParser(p); err != nil {
			log.Fatal(err)
			return c.JSON(fiber.Map{"status": fiber.StatusBadRequest, "message": "Bad Request"})
		}
		randString, err := randomString.Generate(10)
		if err != nil {
			log.Fatal("Error generating random string", err)
			return c.JSON(fiber.Map{ "status": fiber.StatusInternalServerError, "message": "Error generating random string" })
		}
		cleanRandString := strings.Replace(randString, "/", "-", -1)
		redisErr := client.Set(ctx, cleanRandString, p.Url, 0).Err()
		if redisErr != nil {
			log.Fatal("Error storing value in store:\n", redisErr)
			return c.JSON(fiber.Map{"status": fiber.StatusInternalServerError, "message": "Error setting value to store"})
		}
		data := map[string]string{
			"url": "localhost:8000/" + cleanRandString,
		}
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"status":  fiber.StatusCreated,
			"message": "short url created",
			"data":    data,
		})
		
	})

	app.Get("/:url", func(c *fiber.Ctx) error {
		redirectKey := c.Params("url")
		
		// Check if key exists in Redis
		exists, err := client.Exists(ctx, redirectKey).Result()
		if err != nil {
			log.Printf("Error checking Redis key: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status": fiber.StatusInternalServerError,
				"message": "Error checking URL",
			})
		}
		
		if exists == 0 {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status": fiber.StatusNotFound,
				"message": "Short URL not found",
			})
		}

		originalUrl := client.Get(ctx, redirectKey).Val()
		if originalUrl == "" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status": fiber.StatusNotFound,
				"message": "URL not found",
			})
		}

		return c.Redirect(originalUrl, fiber.StatusMovedPermanently)
	})
	app.Listen(":8000")
}