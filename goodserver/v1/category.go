package v1

import "time"

// Category
type Category struct {

	Id int `gorm:"column:id"`

	Name string `gorm:"column:name"`

	ParentCategoryId int `gorm:"column:parent_category_id"`

	Level int `gorm:"column:level"`

	IsTab uint8 `gorm:"column:is_tab"`

	Url string `gorm:"column:url"`

	AddTime time.Time `gorm:"column:add_time"`

	IsDeleted uint8 `gorm:"column:is_deleted"`

	UpdateTime time.Time `gorm:"column:update_time"`


}
