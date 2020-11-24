package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/teten-nugraha/golang-crud-injection/product"
)

func initDB() *gorm.DB{
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang_crud")
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&product.Product{})

	return db
}

func main() {
	db := initDB()
	defer db.Close()

	productAPI := initProductAPI(db)

	r := gin.Default()

	r.GET("/products", productAPI.FindAll)
	r.GET("/products/:id", productAPI.FindById)
	r.POST("/products", productAPI.Create)
	r.PUT("/products/:id", productAPI.Update)
	r.DELETE("/products/:id", productAPI.Delete)

	err := r.Run()
	if err != nil {
		panic(err)
	}
}
