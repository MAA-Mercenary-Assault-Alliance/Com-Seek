package models

import "time"

type File struct {
	ID        string `gorm:"type:uuid;primarykey;" json:"id"`
	UserID    uint
	Extension string
	CreatedAt time.Time
	UpdatedAt time.Time
}
