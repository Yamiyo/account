package po

import (
	"time"
)

// LockPO ...
type LockPO struct {
	Resource   string     `gorm:"column:resource;primary_key"`
	UUID       string     `gorm:"column:uuid"`
	CreateTime *time.Time `gorm:"column:create_time"`
	UpdateTime *time.Time `gorm:"column:update_time"`
	Watch      chan int64 `gorm:"-"`
	CR         chan int64 `gorm:"-"`
	CE         chan error `gorm:"-"`
}

// TableName ...
func (LockPO) TableName() string {
	return "lock"
}
