package item

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"store/model/item"
	"store/util"
)

var db *gorm.DB

func InitCategory(gormdb *gorm.DB, r *gin.Engine) {
	db = gormdb
	db.AutoMigrate(&item.Category{})

	r.POST("/categories", createCategory)
	r.GET("/categories", getAllCategories)
	r.GET("/categories/:id", getCategory)
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

func getAllCategories(c *gin.Context) {
	var categories []item.Category

	token := util.VerifyToken(c.Request.Header.Get("Authorization"))

	if token == nil {
		log.Print("not authroize")
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "not authorize",
		})
		return
	}

	db.Find(&categories)

	c.JSON(http.StatusOK, gin.H{
		"message":    "success",
		"categories": categories,
	})
}

func getCategory(c *gin.Context) {
	var category item.Category
	categoryId, _ := strconv.Atoi(c.Param("id"))

	token := util.VerifyToken(c.Request.Header.Get("Authorization"))

	if token == nil {
		log.Print("not authroize")
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "not authorize",
		})
		return
	}

	db.First(&category, categoryId)

	c.JSON(http.StatusCreated, gin.H{
		"message":  "success",
		"category": category,
	})
}
