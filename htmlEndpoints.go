package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

var ErrorTemplateName = "partials/error"

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

	app.Get("/products", func(ctx *fiber.Ctx) error {
		products := GetAllProducts(db)

		return ctx.Render("pages/products", fiber.Map{
			"Products": products,
		})
	})

	app.Get("/partials/product-edit", func(ctx *fiber.Ctx) error {
		id := ctx.QueryInt("id", 0)

		if id == 0 {
			return ctx.Render(ErrorTemplateName, nil, "")
		}

		byId, err := GetProductById(db, id)
		if err != nil {
			return ctx.Render(ErrorTemplateName, nil, "")
		}

		return ctx.Render("partials/product-edit", fiber.Map{
			"Id":    byId.Id,
			"Name":  byId.Name,
			"Price": byId.Price,
		}, "")
	})

	app.Get("/partials/product", func(ctx *fiber.Ctx) error {
		id := ctx.QueryInt("id", 0)

		if id == 0 {
			return ctx.Render(ErrorTemplateName, nil, "")
		}

		byId, err := GetProductById(db, id)
		if err != nil {
			return ctx.Render(ErrorTemplateName, nil, "")
		}

		return ctx.Render("partials/product", fiber.Map{
			"Id":    byId.Id,
			"Name":  byId.Name,
			"Price": byId.Price,
		}, "")
	})
}
