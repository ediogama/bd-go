package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:products_categories;"`
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	Categories   []Category `gorm:"many2many:products_categories;"`
	SerialNumber SerialNumber
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	category := Category{Name: "Cozinha"}
	db.Create(&category)

	category2 := Category{Name: "Eletronicos"}
	db.Create(&category)

	db.Create(&Product{
		Name:       "Panela Eletrica",
		Price:      100.00,
		Categories: []Category{category, category2},
	})

	db.Create(&SerialNumber{
		Number:    "123456",
		ProductID: 1,
	})

	var categories []Category
	err = db.Model(&Category{}).Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			fmt.Println("- ", product.Name, ", Serial Number:", product.SerialNumber.Number)
		}
	}

}
