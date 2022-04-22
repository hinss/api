package v1

import (
	metav1 "github.com/hinss/api/pkg/meta/v1"
)


type Brand struct {
	metav1.ObjectMeta

	Name string `json:"name" validate:"required" gorm:"column:name;not null"`
	Logo string `json:"logo" gorm:"column:logo"`
}

func (br *Brand) TableName() string {
	return "brands"
}
