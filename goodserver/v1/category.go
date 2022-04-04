package v1

import (
	"time"
	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
)

// Category
type Category struct {
	Id int `gorm:"column:id;primaryKey"`

	Name string `json:"name" gorm:"column:name;not null" validate:"required,min=3,max=20"`

	ParentCategoryId int `json:"parent" gorm:"column:parent_category_id"`

	Level int `json:"level" gorm:"column:level;not null" validate:"required,oneof=1 2 3"`

	IsTab uint8 `json:"is_tab" gorm:"column:is_tab;not null"`

	Url string `gorm:"column:url;not null"`

	AddTime time.Time `gorm:"column:add_time"`

	IsDeleted uint8 `gorm:"column:is_deleted"`

	UpdateTime time.Time `gorm:"column:update_time"`
}

func (c Category) TableName() string {
	return "category"
}

type CategoryList struct {

	metav1.ListMeta

	Items []*Category

}
