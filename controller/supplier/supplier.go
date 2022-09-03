package supplier

import (
	"log"
	"net/http"
	"store/model/supplier"
	"store/util"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func Init(gormdb *gorm.DB, r *gin.Engine) {
	db = gormdb
	db.AutoMigrate(&supplier.Supplier{})

	r.POST("/supplier", addSupplier)
	r.GET("/supplier/:id", getSupplier)
	r.GET("/supplier", getAllSupplier)
	r.PUT("/supplier", updateSupplier)
}

func addSupplier(c *gin.Context) {
	var supplier supplier.Supplier

	token := util.VerifyToken(c.Request.Header.Get("Authorization"))

	if token == nil && token["role"] != "admin" {
		log.Print("not authroize")
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "not authorize",
		})
		return
	}

	if err := c.BindJSON(&supplier); err != nil {
		log.Print("cannot bind supplier")
		return
	}

	db.Create(&supplier)

	c.JSON(http.StatusOK, gin.H{
		"supplier": supplier,
	})
}

func getSupplier(c *gin.Context) {
	var supplier supplier.Supplier
	supplierId, _ := strconv.Atoi(c.Param("id"))

	token := util.VerifyToken(c.Request.Header.Get("Authorization"))

	if token == nil && token["role"] != "admin" {
		log.Print("not authroize")
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "not authorize",
		})
		return
	}

	db.First(&supplier, supplierId)

	c.JSON(http.StatusOK, gin.H{
		"supplier": supplier,
	})
}

func getAllSupplier(c *gin.Context) {
	var suppliers []supplier.Supplier

	token := util.VerifyToken(c.Request.Header.Get("Authorization"))

	if token == nil && token["role"] != "admin" {
		log.Print("not authroize")
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "not authorize",
		})
		return
	}

	db.Find(&suppliers)

	c.JSON(http.StatusOK, gin.H{
		"supplier": suppliers,
	})
}

func updateSupplier(c *gin.Context) {
	var supplier supplier.Supplier

	token := util.VerifyToken(c.Request.Header.Get("Authorization"))

	if token == nil && token["role"] != "admin" {
		log.Print("not authroize")
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "not authorize",
		})
		return
	}

	if err := c.BindJSON(&supplier); err != nil {
		log.Print("cannot bind supplier")
		return
	}

	db.Save(&supplier)

	c.JSON(http.StatusOK, gin.H{
		"supplier": supplier,
	})
}
