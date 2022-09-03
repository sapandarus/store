package user

import (
	"github.com/jinzhu/gorm"
)

func init() {}

type User struct {
	gorm.Model
	StoreId  string `json:"storeId"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
