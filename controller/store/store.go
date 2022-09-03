package store

import (
	"log"
	"net/http"
	"store/model/store"
	"store/util"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func Init(gormdb *gorm.DB, r *gin.Engine) {
	db = gormdb
	db.AutoMigrate(&store.Store{})

	r.POST("/store", addStore)
	r.PATCH("/store", updateStore)
	r.GET("/store", getStore)
}

func addStore(c *gin.Context) {
	var store store.Store

	token := util.VerifyToken(c.Request.Header.Get("Authorization"))

	if token == nil && token["role"] != "admin" {
		log.Print("not authroize")
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "not authorize",
		})
		return
	}

	if err := c.BindJSON(&store); err != nil {
		log.Print("cannot bind store")
		return
	}

	db.Create(&store)

	c.JSON(http.StatusOK, gin.H{
		"store": store,
	})
}

func updateStore(c *gin.Context) {
	var store store.Store

	token := util.VerifyToken(c.Request.Header.Get("Authorization"))

	if token == nil && token["role"] != "admin" {
		log.Print("not authroize")
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "not authorize",
		})
		return
	}

	if err := c.BindJSON(&store); err != nil {
		log.Print("cannot bind store")
		return
	}

	db.Save(&store)

	c.JSON(http.StatusOK, gin.H{
		"store": store,
	})
}

func getStore(c *gin.Context) {
	var store store.Store

	token := util.VerifyToken(c.Request.Header.Get("Authorization"))

	if token == nil {
		log.Print("not authroize")
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "not authorize",
		})
		return
	}

	storeId := token["storeId"]

	db.First(&store, storeId)

	c.JSON(http.StatusOK, gin.H{
		"store": store,
	})
}
