package postgres

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	ProjectID uint64 `gorm:"index:project_id,not null"`
	Data      string `gorm:"size:4096"`
}
