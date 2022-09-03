package store

import "github.com/jinzhu/gorm"

func init() {}

type Store struct {
	gorm.Model
	Address string `json:"address"`
	Phone   string `json:"phone"`
}
