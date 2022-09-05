package item

import "github.com/jinzhu/gorm"

func item() {}

type Item struct {
	gorm.Model
	CategoryId  uint   `json:"categoryId"`
	Name        string `json:"name"`
	Price       uint   `json:"price"`
	Weight      uint   `json:"weight"`
	Unit        string `json:"unit"`
	Description string `json:"description"`
}
