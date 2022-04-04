package v1

import (
	"time"
	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
)

// Category
type Category struct {
	Id int `gorm:"column:id;primaryKey"`

	Name string `gorm:"column:name;not null"`

	ParentCategoryId int `gorm:"column:parent_category_id"`

	Level int `gorm:"column:level"`

	IsTab uint8 `gorm:"column:is_tab;not null"`

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
