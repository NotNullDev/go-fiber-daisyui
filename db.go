package main

import (
	"errors"
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

func GetAllProducts(db *sqlx.DB) []Product {
	var products []Product

	err := db.Select(&products, "SELECT * FROM products")

	if err != nil {
		log.Fatalln(err)
	}

	return products
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
