package item

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"store/model/item"
	"store/util"
)

var db *gorm.DB

func Init(gormdb *gorm.DB, r *gin.Engine) {
	db = gormdb
	db.AutoMigrate(&item.Category{})

	r.POST("/categories", createCategory)
}

func createCategory(c *gin.Context) {
	var category item.Category

	token := util.VerifyToken(c.Request.Header.Get("Authorization"))

	if token == nil && token["role"] != "admin" {
		log.Print("not authroize")
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "not authorize",
		})
		return
	}

	if err := c.BindJSON(&category); err != nil {
		log.Print("cannot bind category")
		return
	}

	db.Create(&category)

	c.JSON(http.StatusCreated, gin.H{
		"message":  "created",
		"category": category,
	})
}
