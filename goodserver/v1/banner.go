package v1

import (
	metav1 "github.com/hinss/api/pkg/meta/v1"
)

type Banners struct {

	metav1.ObjectMeta

	Name string `gorm:"column:name"`

	Logo string `gorm:"column:logo"`
}
