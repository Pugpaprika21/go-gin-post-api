package models

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	PostID    uint
	Post      Post `gorm:"foreignKey:PostID"`
	UserID    uint
	User      User   `gorm:"foreignKey:UserID"`
	Content   string `gorm:"type:varchar(255); not null"`
	Timestamp time.Time
}
