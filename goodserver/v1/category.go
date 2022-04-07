package v1

import (
	metav1 "github.com/marmotedu/component-base/pkg/meta/v1"
	"time"
)

// Category
type Category struct {
	Id int `json:"id" gorm:"column:id;primaryKey"`

	Name string `json:"name" validate:"required,min=3,max=20" gorm:"column:name;not null" `

	ParentCategoryId *int `json:"parent" validate:"omitempty" gorm:"column:parent_category_id"`

	Level int `json:"level" validate:"required,oneof=1 2 3" gorm:"column:level;not null" `

	IsTab *bool `json:"is_tab" validate:"required" gorm:"column:is_tab;not null"`

	Url string `gorm:"column:url;not null"`

	AddTime time.Time `gorm:"column:add_time"`

	IsDeleted uint8 `gorm:"column:is_deleted"`

	UpdateTime time.Time `gorm:"column:update_time"`
}

func (c *Category) TableName() string {
	return "category"
}

func (c *Category) NewCategoryInfo() *CategoryInfo{
	return &CategoryInfo{
		Id: c.Id,
		Name: c.Name,
		SubCategorys: make([]*CategoryInfo, 0),
	}
}

type CategoryList struct {
	metav1.ListMeta

	Items []*Category

	jsonData string
}

type CategoryInfo struct {
	Id int `json:"id"`

	Name string `json:"name"`

	SubCategorys []*CategoryInfo `json:"sub_category"`
}
