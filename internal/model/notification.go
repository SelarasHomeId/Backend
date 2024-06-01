package model

import (
	"selarashomeid/internal/abstraction"
	"time"
)

type NotificationEntity struct {
	Title     string    `json:"title"`
	Message   string    `json:"message"`
	IsRead    bool      `json:"is_read"`
	CreatedAt time.Time `json:"created_at"`
}

// NotificationEntityModel ...
type NotificationEntityModel struct {
	ID int `json:"id" param:"id" form:"id" validate:"number,min=1" gorm:"primaryKey;autoIncrement;"`

	// entity
	NotificationEntity

	// context
	Context *abstraction.Context `json:"-" gorm:"-"`
}

// TableName ...
func (NotificationEntityModel) TableName() string {
	return "notification"
}
