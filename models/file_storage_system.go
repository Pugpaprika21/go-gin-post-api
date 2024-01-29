package models

import "gorm.io/gorm"

type FileStorageSystem struct {
	gorm.Model
	FileName      string `gorm:"type:varchar(255)"`
	FileSize      int64
	FileType      string `gorm:"type:varchar(50)"`
	FileExtension string `gorm:"type:varchar(50)"`
	Content       string `gorm:"type:varchar(255)"`
	RefID         uint
	RefTable      string `gorm:"type:varchar(100)"`
	RefField      string `gorm:"type:varchar(100)"`
}
