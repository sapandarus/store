package supplier

import "github.com/jinzhu/gorm"

func init() {}

type Supplier struct {
	gorm.Model
	Brand   string `json:"brand"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}
