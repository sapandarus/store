package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"store/controller/item"
	"store/controller/store"
	"store/controller/supplier"
	"store/controller/user"
)

func main() {
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/store?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	defer db.Close()

	r := gin.Default()

	user.Init(db, r)
	store.Init(db, r)
	supplier.Init(db, r)
	item.InitCategory(db, r)
	item.InitItemInformation(db, r)

	r.Run(":8080")
}
