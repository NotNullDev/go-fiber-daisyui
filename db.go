package main

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Product struct {
	Id    int    `db:"id"`
	Name  string `db:"name"`
	Price int    `db:"price"`
}

func NewDb() *sqlx.DB {
	db, err := sqlx.Connect("sqlite3", "./db.sqlite")

	if err != nil {
		log.Fatalln(err)
	}

	return db
}

func AddProduct(db *sqlx.DB, name string, price int) (int, error) {

	if strings.TrimSpace(name) == "" {
		return 0, errors.New("name is empty")
	}

	if price == 0 {
		return 0, errors.New("price must be greater than 0")
	}

	cursor, err := db.Exec("INSERT INTO products (name, price) VALUES (?, ?)", name, price)
	if err != nil {
		return 0, err
	}

	id, err := cursor.LastInsertId()

	return int(id), err
}

func UpdateProduct(db *sqlx.DB, id int, name string, price int) error {
	if strings.TrimSpace(name) == "" {
		return errors.New("name is empty")
	}

	if price == 0 {
		return errors.New("price must be greater than 0")
	}

	_, err := db.Exec("UPDATE products SET name = ?, price = ? WHERE id = ?", name, price, id)

	return err
}

func GetAllProducts(db *sqlx.DB) []Product {
	var products []Product

	err := db.Select(&products, "SELECT * FROM products")

	if err != nil {
		log.Fatalln(err)
	}

	return products
}

func GetFilteredProducts(db *sqlx.DB, filter string) ([]Product, error) {
	var products []Product
	filterParsed := fmt.Sprintf("%%%s%%", strings.TrimSpace(strings.ToLower(filter)))
	err := db.Select(&products, "select * from products p where trim(lower(p.name)) like ?", filterParsed)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func GetProductById(db *sqlx.DB, id int) (Product, error) {
	var product Product

	err := db.Get(&product, "SELECT * FROM products WHERE id = ?", id)

	return product, err
}

func deleteProduct(db *sqlx.DB, id int) error {
	_, err := db.Exec("delete from products where id = ?", id)
	return err
}
