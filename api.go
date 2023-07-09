package main

import (
	"errors"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type CreateProductRequest struct {
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
		var product CreateProductRequest
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

	apiGroup.Put("/products", func(c *fiber.Ctx) error {
		var product Product
		if err := c.BodyParser(&product); err != nil {
			return err
		}

		err := UpdateProduct(db, product.Id, product.Name, product.Price)
		if err != nil {
			return err
		}

		return c.Render("partials/product", fiber.Map{
			"Id":    product.Id,
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
}
