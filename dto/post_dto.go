package dto

import "github.com/gin-gonic/gin"

type PostBodyRequest struct {
	UserID  uint   `form:"userId" binding:"required"`
	Title   string `form:"title" binding:"required"`
	Content string `form:"content" binding:"required"`
}

type PostResponse struct {
	StatusBool    bool   `json:"StatusBool"`
	StatusMessage string `json:"StatusMessage"`
	Data          gin.H  `json:"Data"`
}

type PostData struct {
	ID      uint   `gorm:"column:id"`
	UserID  uint   `gorm:"column:user_id"`
	Title   string `gorm:"column:title"`
	Content string `gorm:"column:content"`
}

type PostFetchFileData struct {
	ID       uint   `gorm:"column:id"`
	Filename string `gorm:"column:file_name"`
}

type PostFetchData struct {
	ID      uint   `gorm:"column:id"`
	UserID  uint   `gorm:"column:user_id"`
	Title   string `gorm:"column:title"`
	Content string `gorm:"column:content"`
}
