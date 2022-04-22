package v1

import (
	metav1 "github.com/hinss/api/pkg/meta/v1"
)

// Category
type Category struct {

	metav1.ObjectMeta

	Name string `json:"name" validate:"required,min=3,max=20" gorm:"column:name;not null" `

	ParentCategoryId *int `json:"parent" validate:"omitempty" gorm:"column:parent_category_id"`

	Level int `json:"level" validate:"required,oneof=1 2 3" gorm:"column:level;not null" `

	IsTab *bool `json:"is_tab" validate:"required" gorm:"column:is_tab;not null"`

	Url string `json:"url" gorm:"column:url;not null"`

}

func (c *Category) TableName() string {
	return "category"
}

func (c *Category) NewCategoryInfo() *CategoryInfo {
	return &CategoryInfo{
		Category:     c,
		SubCategorys: make([]*CategoryInfo, 0),
	}
}

type CategoryList struct {
	metav1.ListMeta

	Items []*Category

	JsonData string
}

type CategoryInfo struct {
	*Category

	SubCategorys []*CategoryInfo `json:"sub_category"`
}

// CategoryLevelInfo use to inner join second level and third level categroy
type CategoryLevelInfo struct {
	Pid         int    `gorm:"column:pId"`
	Pname       string `gorm:"column:pName"`
	ParentId    int    `gorm:"column:parentId"`
	Plevel      int    `gorm:"column:pLevel"`
	SubId       int    `gorm:"column:subId"`
	SubName     string `gorm:"column:subName"`
	SubParentId int    `gorm:"column:subPid"`
	SubLevel    int    `gorm:"column:subLevel"`
}
