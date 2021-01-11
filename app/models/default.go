package models

import (
	"time"
)

type DefaultModel struct {
	ID        uint      `gorm:"column:id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	RemovedAt time.Time `gorm:"column:removed_at" json:"removed_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}
