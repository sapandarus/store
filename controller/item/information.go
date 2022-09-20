package item

import (
	"log"
	"net/http"
	"store/model/item"
	"store/util"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func InitItemInformation(gormdb *gorm.DB, r *gin.Engine) {
	db = gormdb
	db.AutoMigrate(&item.Item{})

	r.POST("/items", createItem)
	r.GET("/items", getItems)
}

func createItem(c *gin.Context) {
	var item item.Item

	token := util.VerifyToken(c.Request.Header.Get("Authorization"))

	if token == nil && token["role"] != "admin" {
		log.Print("not authroize")
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "not authorize",
		})
		return
	}

	if err := c.BindJSON(&item); err != nil {
		log.Print("cannot bind item information")
		return
	}

	db.Create(&item)

	c.JSON(http.StatusCreated, gin.H{
		"message": "created",
		"item_information": item,
	})
}

func getItems(c *gin.Context) {
	var itemRequest item.ItemRequest
	var items []item.Item

	token := util.VerifyToken(c.Request.Header.Get("Authorization"))

	if token == nil {
		log.Print("not authroize")
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "not authorize",
		})
		return
	}

	if err := c.ShouldBind(&itemRequest); err != nil {
		log.Print("cannot bind item request")
		return
	}

	baseQuery := db
	
	if itemRequest.Id > 0 {
		baseQuery = baseQuery.Where(&item.Item{Model: gorm.Model{ID: itemRequest.Id}})
	}

	if itemRequest.CategoryId > 0 {
		baseQuery = baseQuery.Where(&item.Item{CategoryId: itemRequest.CategoryId})
	}

	baseQuery.Find(&items)

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"items": items,
	})
}