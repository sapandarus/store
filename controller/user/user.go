package user

import (
	"log"
	"net/http"
	"store/model/user"
	"store/util"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func Init(gormdb *gorm.DB, r *gin.Engine) {
	db = gormdb
	db.AutoMigrate(&user.User{})

	r.POST("/login", login)
	r.POST("/users", createUser)
	r.GET("/users", getAllUser)
	r.GET("/users/:id", getUser)
}

func login(c *gin.Context) {
	var login user.Login
	var user user.User

	if err := c.BindJSON(&login); err != nil {
		log.Print("connot bind login")
		return
	}

	db.First(&user, "username = ? AND password = ?", login.Username, login.Password)

	if user.Username != "" {

		tokenString := util.CreateToken(user)

		c.JSON(http.StatusOK, gin.H{
			"token": tokenString,
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "user not found",
		})
	}
}

func createUser(c *gin.Context) {
	var user user.User

	if err := c.BindJSON(&user); err != nil {
		log.Print("connot bind create user")
		return
	}

	token := util.VerifyToken(c.Request.Header.Get("Authorization"))

	if token == nil && token["role"] != "admin" {
		log.Print("not authroize")
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "not authorize",
		})
		return
	}

	db.Create(&user)

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func getUser(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	var user user.User

	token := util.VerifyToken(c.Request.Header.Get("Authorization"))

	if token == nil {
		log.Print("not authroize")
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "not authorize",
		})
		return
	}

	db.Select([]string{"id", "store_id", "name", "username", "role"}).First(&user, userId)

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func getAllUser(c *gin.Context) {
	var users []user.User

	token := util.VerifyToken(c.Request.Header.Get("Authorization"))

	if token == nil {
		log.Print("not authroize")
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "not authorize",
		})
		return
	}

	db.Select([]string{"id", "store_id", "name", "username", "role"}).Find(&users)

	c.JSON(http.StatusOK, gin.H{
		"user": users,
	})
}
