package repository

import (
	"time"

	"github.com/Pugpaprika21/go-gin/db"
	"github.com/Pugpaprika21/go-gin/dto"
	"github.com/Pugpaprika21/go-gin/models"
	"gorm.io/gorm"
)

type CommentRepositoryInterface interface {
	CreateUserComment(body dto.CommentBodyRequest) bool
	FindAllComment() []dto.CommentFetchData
	FindCommentByID(commentID uint) dto.CommentFetchData
	FindCommentByUserID(userID uint) []dto.CommentFetchData
	UpdateCommentByID(commentID uint, body dto.CommentBodyRequest) bool
	DeleteCommentByID(commentID uint) bool
}

type CommentRepository struct {
	DB *gorm.DB
}

func NewCommentRepository() *CommentRepository {
	return &CommentRepository{
		DB: db.Conn,
	}
}

func (c *CommentRepository) CreateUserComment(body dto.CommentBodyRequest) bool {
	var post dto.PostFetchData
	if c.DB.Model(&models.Post{}).Where("id = ?", body.PostID).Find(&post).RowsAffected == 0 {
		return false
	}
	comment := models.Comment{
		PostID:  body.PostID,
		UserID:  body.UserID,
		Content: body.Content,
	}
	return c.DB.Create(&comment).Error == nil
}

func (c *CommentRepository) FindAllComment() []dto.CommentFetchData {
	var comments []dto.CommentFetchData
	c.DB.Model(&models.Comment{}).Order("created_at DESC").Find(&comments)
	return comments
}

func (c *CommentRepository) FindCommentByID(commentID uint) dto.CommentFetchData {
	var comment dto.CommentFetchData
	c.DB.Model(&models.Comment{}).Where("id = ?", commentID).Find(&comment)
	return comment
}

func (c *CommentRepository) FindCommentByUserID(userID uint) []dto.CommentFetchData {
	var comments []dto.CommentFetchData
	c.DB.Model(&models.Comment{}).Where("user_id = ?", userID).Order("created_at DESC").Find(&comments)
	return comments
}

func (c *CommentRepository) UpdateCommentByID(commentID uint, body dto.CommentBodyRequest) bool {
	var comment dto.CommentFetchData
	if c.DB.Model(&models.Comment{}).Where("id = ?", commentID).Find(&comment).RowsAffected == 0 {
		return false
	}

	updateComment := map[string]interface{}{
		"PostID":    body.PostID,
		"UserID":    body.UserID,
		"Content":   body.Content,
		"UpdatedAt": time.Now(),
	}
	return c.DB.Model(&models.Comment{}).Where("id = ?", commentID).Updates(updateComment).Error == nil
}

func (c *CommentRepository) DeleteCommentByID(commentID uint) bool {
	var comment dto.CommentFetchData
	if c.DB.Model(&models.Comment{}).Where("id = ?", commentID).Find(&comment).RowsAffected == 0 {
		return false
	}
	return c.DB​.Delete(&models.Comment{}, commentID).Error == nil
}
