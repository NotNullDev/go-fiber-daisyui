package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func initHtmlEndpoints(app *fiber.App, db *sqlx.DB) {
	app.Static("/", "./public")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("pages/main", fiber.Map{
			"Title": "Hello, World!",
		})
	})

	app.Get("/about", func(c *fiber.Ctx) error {
		return c.Render("pages/about", fiber.Map{
			"Title": "Hello, World!",
		})
	})
}
