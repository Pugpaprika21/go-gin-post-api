package repository

import (
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/Pugpaprika21/go-gin/db"
	"github.com/Pugpaprika21/go-gin/dto"
	"github.com/Pugpaprika21/go-gin/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostRepositoryInterface interface {
	GetPosts(postID ...uint) []dto.PostFetchData
	GetFiles(postID uint) []dto.PostFetchFileData
	GetComments(postID uint) []dto.CommentFetchData
	FindPostAll(postID ...uint) []map[string]interface{}
	DeletePostByID(postID uint) bool
	GetPostByUserID(userID uint) []dto.PostFetchData
	CreatePostByUser(body dto.PostBodyRequest, files []*multipart.FileHeader, ctx *gin.Context) bool
	uploadPostAttachments(postID uint, files []*multipart.FileHeader, ctx *gin.Context, action string)
	UpdatePostByID(postID uint, body dto.PostBodyRequest, files []*multipart.FileHeader, ctx *gin.Context) bool
}

type PostRepository struct {
	DB *gorm.DB
}

func NewPostRepository() *PostRepository {
	return &PostRepository{
		DB: db.Conn,
	}
}

func (p *PostRepository) GetPosts(postID ...uint) []dto.PostFetchData {
	var posts []dto.PostFetchData
	query := p.DB.Model(&models.Post{}).Select("id", "title", "content")
	if len(postID) > 0 {
		query.Where("id = ?", postID[0])
	}
	query.Find(&posts)
	return posts
}

func (p *PostRepository) GetFiles(postID uint) []dto.PostFetchFileData {
	var files []dto.PostFetchFileData
	p.DB.Model(&models.FileStorageSystem{}).Select("id", "file_name").Where("ref_id = ? AND ref_table = ?", postID, "post").Find(&files)
	return files
}

func (p *PostRepository) GetComments(postID uint) []dto.CommentFetchData {
	var comments []dto.CommentFetchData
	p.DB.Model(&models.Comment{}).Where("post_id = ?", postID).Find(&comments)
	return comments
}

func (p *PostRepository) FindPostAll(postID ...uint) []map[string]interface{} {
	var postAll []map[string]interface{}
	var posts []dto.PostFetchData

	if len(postID) > 0 {
		posts = p.GetPosts(postID[0])
	} else {
		posts = p.GetPosts()
	}

	for _, post := range posts {
		var filesData []map[string]interface{}
		var commentsData []map[string]interface{}
		for _, file := range p.GetFiles(post.ID) {
			fileData := map[string]interface{}{
				"id":       file.ID,
				"filename": os.Getenv("APP_URL") + "assets/uploads/" + file.Filename,
			}
			filesData = append(filesData, fileData)
		}

		for _, comment := range p.GetComments(post.ID) {
			commentData := map[string]interface{}{
				"id":      comment.ID,
				"content": comment.Content,
			}
			commentsData = append(commentsData, commentData)
		}

		postData := map[string]interface{}{
			"postID":   post.ID,
			"title":    post.Title,
			"content":  post.Content,
			"files":    filesData,
			"comments": commentsData,
		}
		postAll = append(postAll, postData)
	}
	return postAll
}

func (p *PostRepository) CreatePostByUser(body dto.PostBodyRequest, files []*multipart.FileHeader, ctx *gin.Context) bool {
	post := models.Post{
		UserID:    body.UserID,
		Title:     body.Title,
		Content:   body.Content,
		Timestamp: time.Now(),
	}
	if err := p.DB.Create(&post).Error; err == nil {
		p.uploadPostAttachments(post.ID, files, ctx, "create")
		return true
	}
	return false
}

func (p *PostRepository) uploadPostAttachments(postID uint, files []*multipart.FileHeader, ctx *gin.Context, action string) {
	if len(files) == 0 {
		return
	}

	for _, file := range files {
		fileExtension := filepath.Ext(file.Filename)
		fileName := uuid.New().String() + fileExtension

		fileSystem := models.FileStorageSystem{
			FileName:      fileName,
			FileSize:      file.Size,
			FileExtension: fileExtension,
			RefID:         postID,
			RefTable:      "post",
			RefField:      "files",
		}

		ctx.SaveUploadedFile(file, "./assets/uploads/"+fileName)

		if action == "create" {
			p.DB.Create(&fileSystem)
		} else if action == "update" {
			updateFile := map[string]interface{}{
				"FileName":      fileName,
				"FileSize":      file.Size,
				"FileType":      "",
				"FileExtension": fileExtension,
				"UpdatedAt":     time.Now(),
			}
			p.DB.Model(&models.FileStorageSystem{}).Where("ref_id = ? AND ref_table = ?", postID, "post").Updates(updateFile)
		}
	}
}

func (p *PostRepository) GetPostByUserID(userID uint) []dto.PostFetchData {
	var posts []dto.PostFetchData
	p.DB.Model(&models.Post{}).Select("id", "user_id", "title", "content").Where("user_id = ?", userID).Find(&posts)
	return posts
}

func (p *PostRepository) UpdatePostByID(postID uint, body dto.PostBodyRequest, files []*multipart.FileHeader, ctx *gin.Context) bool {
	post := map[string]interface{}{
		"Title":     body.Title,
		"Content":   body.Content,
		"UpdatedAt": time.Now(),
	}

	if err := p.DB.Model(&models.Post{}).Where("id = ?", postID).Updates(post).Error; err == nil {
		p.uploadPostAttachments(postID, files, ctx, "update")
		return true
	}
	return false
}

func (p *PostRepository) DeletePostByID(postID uint) bool {
	var post dto.PostFetchData
	var files models.FileStorageSystem

	if p.DB.Model(&models.Post{}).Where("id = ?", postID).Find(&post).RowsAffected == 0 {
		return false
	}

	for _, file := range p.GetFiles(postID) {
		p.DB.Delete(&files, file.ID)
	}
	return p.DB.Delete(&models.Post{}, postID).Error == nil
}
