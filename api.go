package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func initApiEndpoints(app *fiber.App, db *sqlx.DB) {
	apiGroup := app.Group("/api")

	apiGroup.Get("/data", func(c *fiber.Ctx) error {
		time.Sleep(3 * time.Second)
		return c.JSON(fiber.Map{
			"message": "Hello, World!",
		})
	})
}
