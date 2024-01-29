package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	UserID    uint
	User      User   `gorm:"foreignKey:UserID"`
	Title     string `gorm:"type:varchar(255); not null"`
	Content   string `gorm:"type:varchar(255); not null"`
	Timestamp time.Time
	Comments  []Comment
	Files     []FileStorageSystem `gorm:"foreignKey:RefID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
