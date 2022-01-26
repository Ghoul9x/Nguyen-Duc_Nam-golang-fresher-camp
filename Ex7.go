package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	Id    int    `json:"id,omitempty" gorm:"column:id;`
	Name  string `json:"name gorm:"column:name;"`
	Image string `json:"image,omitempty gorm:"column:image;"`
}

func (Category) TableName() string {
	return "category"
}

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "food_delivery:123456789@tcp(127.0.0.1:3306)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// insert new category
	// newCate := Category{Name: "Pizza", Image: "ádasd"}
	// db.Create(&newCate)
	// if err := db.Create(&newCate); err != nil {
	// 	fmt.Println(err)
	// }

	//Search
	var categories []Category
	if err := db.Where("Name = ?", "Pizza").Find(&categories); err != nil {
		log.Println(err)
	}
	fmt.Println(categories)

	var category Category
	if err := db.Where("id = 4").Find(&category); err != nil {
		log.Println(err)
	}
	//delete category
	db.Table(Category{}.TableName()).Where("id = 1").Delete(nil)
	fmt.Println(category)
	//update category
	db.Table(Category{}.TableName()).Where("id = 1").Updates(map[string]interface{}{
		"name": "Burger",
	})
}
