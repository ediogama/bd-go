package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	CategoryId int
	Category   Category
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{})

	// categoryEletro := Category{Name: "Eletronicos"}
	// db.Create(&categoryEletro)

	db.Create(&Product{
		Name:       "Mouse",
		Price:      50.00,
		CategoryId: 1,
	})

	var products []Product
	db.Preload("Category").Find(&products)
	for _, p := range products {
		fmt.Println(p.Name, p.Category.Name)
	}

}
