package v1

import (
	metav1 "github.com/hinss/api/pkg/meta/v1"

)

type Banners struct {
	metav1.ObjectMeta

	Image string `json:"image" validate:"required" gorm:"column:name;not null"`

	Url string `json:"url" validate:"required" gorm:"column:logo;not null"`

	Index int `json:"index" validate:"required" gorm:"column:name;not null"`
}

func (b *Banners) TableName() string {
	return "banner"
}

