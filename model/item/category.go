package item

import "github.com/jinzhu/gorm"

func init() {}

type Category struct {
	gorm.Model
	ParentId uint   `json:"parentId"`
	Name     string `json:"name"`
}
