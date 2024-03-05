package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

// loggingMiddleware logs the processing time for each request
func loggingMiddleware(c *fiber.Ctx) error {
	// Start timer
	start := time.Now()

	// Process request
	err := c.Next()

	// Calculate processing time
	duration := time.Since(start)

	// Log the information
	fmt.Printf("Request URL: %s - Method: %s - Duration: %s\n", c.OriginalURL(), c.Method(), duration)

	return err
}

func main() {
	app := fiber.New()

	// Use the logging middleware
	app.Use(loggingMiddleware)

	// Setup routes
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World! middleware")
	})

	app.Listen(":8080")
}
