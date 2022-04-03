package v1

import "time"

type Banners struct {
	Id int `gorm:"column:id"`

	Name string `gorm:"column:name"`

	Logo string `gorm:"column:logo"`

	AddTime time.Time `gorm:"column:add_time"`

	IsDeleted uint8 `gorm:"column:is_deleted"`

	UpdateTime time.Time `gorm:"column:update_time"`
}
