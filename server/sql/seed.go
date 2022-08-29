package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Product struct {
	Id   int64  `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Name string `gorm:"column:name;type:varchar(255);not null"`
}

type Category struct {
	Id    int64  `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Label string `gorm:"column:label;type:varchar(255);not null"`
}

type Order struct {
	Id     int64  `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	UserId string `gorm:"column:user_id;type:int;not null`
}

type OrderProduct struct {
	Id        int64  `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	OrderId   string `gorm:"column:order_id;type:int;not null;references:order(id)"`
	ProductId string `gorm:"column:product_id;type:int;not null;references:product(id)"`
	Quantity  int64  `gorm:"column:quantity;type:int;not null"`
}

var products = []Product{
	{Id: 1, Name: "Prod1"},
	{Id: 2, Name: "Prod2"},
	{Id: 3, Name: "Prod3"},
	{Id: 4, Name: "Prod4"},
}

var categories = []Category{
	{Id: 1, Label: "Cat1"},
	{Id: 2, Label: "Cat2"},
}

func main() {
	db, err := gorm.Open("postgres", "postgres://postgres:docker@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	for _, product := range products {
		db.Create(&product)
		fmt.Println("Product created: ", product)
	}
	for _, category := range categories {
		db.Create(&category)
		fmt.Println("Category created: ", category)
	}
}
