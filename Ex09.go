package main

import (
	"log"

	"net/http"

	"github.com/gin-gonic/gin"
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
		log.Fatalln(err)
	}

	if err := runService(db); err != nil {
		log.Fatalln(err)
	}

	// insert new category
	// newCate := Category{Name: "Pizza", Image: "ádasd"}
	// db.Create(&newCate)
	// if err := db.Create(&newCate); err != nil {
	// 	fmt.Println(err)
	// }

	//Search
	// var categories []Category
	// if err := db.Where("Name = ?", "Pizza").Find(&categories); err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(categories)

	// var category Category
	// if err := db.Where("id = 3").Find(&category); err != nil {
	// 	log.Println(err)
	// }
	// //delete category
	// db.Table(Category{}.TableName()).Where("id = 1").Delete(nil)
	// if err := db.Where("id = 1").Find(&category); err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(categories)
	// //update category
	// db.Table(Category{}.TableName()).Where("id = 2").Updates(map[string]interface{}{
	// 	"name": "Pizza",
	// })
	// if err := db.Where("id = 2").Find(&category); err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(categories)
}

func runService(db *gorm.DB) error {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	return r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
