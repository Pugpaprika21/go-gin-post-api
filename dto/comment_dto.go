package dto

import "github.com/gin-gonic/gin"

type CommentBodyRequest struct {
	PostID  uint   `form:"postId" binding:"required"`
	UserID  uint   `form:"userId" binding:"required"`
	Content string `form:"content" binding:"required"`
}

type CommentResponse struct {
	StatusBool    bool   `json:"StatusBool"`
	StatusMessage string `json:"StatusMessage"`
	Data          gin.H  `json:"Data"`
}

type CommentFetchData struct {
	ID      uint   `gorm:"column:id"`
	PostID  uint   `gorm:"column:post_id"`
	UserID  uint   `gorm:"column:user_id"`
	Content string `gorm:"column:content"`
}
