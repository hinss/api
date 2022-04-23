package v1

import (
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
	"time"
)

// ObjectMeta is metadata that all persisted resources must have, which includes all objects
// ObjectMeta is also used by gorm.
type ObjectMeta struct {
	Id         int                   `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	AddTime    time.Time             `gorm:"column:add_time;not null" json:"add_time"`
	UpdateTime time.Time             `gorm:"column:update_time;not null" json:"update_time"`
	IsDeleted  soft_delete.DeletedAt `json:"isDeleted" gorm:"column:is_deleted;softDelete:flag"`
}

func (o *ObjectMeta) BeforeUpdate(db *gorm.DB) error {
	o.UpdateTime = time.Now()

	return nil
}

func (o *ObjectMeta) BeforeCreate(db *gorm.DB) error {
	o.AddTime = time.Now()
	o.UpdateTime = time.Now()

	return nil
}

type ListMeta struct {
	TotalCount int64 `json:"totalCount,omitempty"`
}
