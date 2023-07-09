package main

import (
	"log"

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

func AddProduct(db *sqlx.DB, name, price string) {
	_, err := db.Exec("INSERT INTO products (name, price) VALUES (?, ?)", name, price)

	if err != nil {
		log.Fatalln(err)
	}
}

func GetAllProducts(db *sqlx.DB) []Product {
	var products []Product

	err := db.Select(&products, "SELECT * FROM products")

	if err != nil {
		log.Fatalln(err)
	}

	return products
}

func deleteProduct(db *sqlx.DB, id int) error {
	_, err := db.Exec("delete from products where id = ?", id)
	return err
}
