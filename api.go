package main

import (
	"errors"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type ModifyProductRequest struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func initApiEndpoints(app *fiber.App, db *sqlx.DB) {
	apiGroup := app.Group("/api")

	apiGroup.Get("/data", func(c *fiber.Ctx) error {
		time.Sleep(3 * time.Second)
		return c.JSON(fiber.Map{
			"message": "Hello, World!",
		})
	})

	apiGroup.Post("/products", func(c *fiber.Ctx) error {
		var product ModifyProductRequest
		if err := c.BodyParser(&product); err != nil {
			return err
		}

		id, err := AddProduct(db, product.Name, product.Price)
		if err != nil {
			return err
		}

		return c.Render("partials/product", fiber.Map{
			"Id":    id,
			"Name":  product.Name,
			"Price": product.Price,
		}, "")
	})

	apiGroup.Put("/products/:id", func(c *fiber.Ctx) error {
		var product ModifyProductRequest
		if err := c.BodyParser(&product); err != nil {
			return err
		}

		id, err := c.ParamsInt("id", 0)
		if err != nil {
			return err
		}

		if id == 0 {
			return errors.New("id must be greater than 0")
		}

		err = UpdateProduct(db, id, product.Name, product.Price)
		if err != nil {
			return err
		}

		return c.Render("partials/product", fiber.Map{
			"Id":    id,
			"Name":  product.Name,
			"Price": product.Price,
		}, "")
	})

	apiGroup.Delete("/products/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		idParsed, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return err
		}

		if idParsed == 0 {
			return errors.New("id must be greater than 0")
		}

		err = deleteProduct(db, int(idParsed))
		if err != nil {
			return err
		}

		return c.SendString("")
	})

	apiGroup.Post("/products/filter", func(ctx *fiber.Ctx) error {
		filter := ctx.FormValue("filter", "")

		var products []Product

		if filter == "" {
			products = GetAllProducts(db)
		}

		products, err := GetFilteredProducts(db, filter)
		if err != nil {
			return err
		}

		return ctx.Render("partials/filter-search-result", fiber.Map{
			"Products": products,
		}, "")
	})
}
