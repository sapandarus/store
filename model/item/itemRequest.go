package item

func init() {}

type ItemRequest struct {
	Id         uint `form:"id"`
	CategoryId uint `form:"categoryId"`
}