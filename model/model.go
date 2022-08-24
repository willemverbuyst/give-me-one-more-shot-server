package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MessageID struct {
	ID string `uri:"id" binding:"required"`
}

type Message struct {
	ID        string         `gorm:"primaryKey" json:"id"`
	Content   string         `json:"message" binding:"required"`
	CreatedAt int64          `gorm:"autoCreateTime:milli" json:"created_at"`
	UpdatedAt int64          `gorm:"autoUpdateTime:milli" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

func (x *Message) FillDefaults() {
	if x.ID == "" {
		x.ID = uuid.New().String()
	}
}
