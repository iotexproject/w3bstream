package coordinator

import (
	"github.com/machinefi/sprout/types"
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	MessageID string             `gorm:"index:message_id,not null"`
	ProjectID uint64             `gorm:"index:message_fetch,not null"`
	Data      string             `gorm:"size:4096"`
	State     types.MessageState `gorm:"index:message_fetch,not null"`
}

type MessageStateLog struct {
	gorm.Model
	MessageID string             `gorm:"index:message_id,not null"`
	State     types.MessageState `gorm:"not null"`
	Comment   string
}
