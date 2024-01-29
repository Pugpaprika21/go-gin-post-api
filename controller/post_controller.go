package controller

import (
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/Pugpaprika21/go-gin/db"
	"github.com/Pugpaprika21/go-gin/dto"
	"github.com/Pugpaprika21/go-gin/models"
	"github.com/Pugpaprika21/go-gin/repository"
	"github.com/gin-gonic/gin"
)

type Post struct {
	repository.PostRepositoryInterface
}

func (p Post) GetPostAll(ctx *gin.Context) {
	postID := ctx.Param("postId")
	postID = strings.TrimPrefix(postID, "/")

	var posts []map[string]interface{}
	if postID != "" {
		postID, _ := strconv.ParseUint(postID, 10, 64)
		posts = p.FindPostAll(uint(postID))
	} else {
		posts = p.FindPostAll()
	}

	if len(posts) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"Data": nil, "Message": "Post Not Found .."})
		return
	}

	ctx.JSON(http.StatusOK, dto.PostResponse{
		StatusBool:    true,
		StatusMessage: "Post..",
		Data:          gin.H{"posts": posts},
	})
}

func (p Post) ShowPostAttachments(ctx *gin.Context) {
	filename := ctx.Param("filename")
	imagePath := "./assets/uploads/" + filename

	imageFile, err := os.Open(imagePath)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
		return
	}
	defer imageFile.Close()

	fileInfo, _ := imageFile.Stat()
	var fileSize int64 = fileInfo.Size()
	buffer := make([]byte, fileSize)

	_, err = imageFile.Read(buffer)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read image"})
		return
	}
	ctx.File(imagePath)
}

func (p Post) GetPostByUser(ctx *gin.Context) {
	userID, _ := strconv.ParseUint(ctx.Param("userId"), 10, 64)

	posts := p.GetPostByUserID(uint(userID))
	if len(posts) == 0 {
		ctx.JSON(http.StatusOK, gin.H{"data": nil, "Message": "Post Not Found ..."})
		return
	}

	ctx.JSON(http.StatusOK, dto.PostResponse{
		StatusBool:    true,
		StatusMessage: "Post By User ..",
		Data: gin.H{
			"Posts": posts,
		},
	})
}

func (p Post) CreatePost(ctx *gin.Context) {
	var body dto.PostBodyRequest
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	form, _ := ctx.MultipartForm()
	files := form.File["files"]

	if !p.CreatePostByUser(body, files, ctx) {
		ctx.JSON(http.StatusOK, gin.H{"Data": nil, "Message": "Create Post Error.. "})
		return
	}

	ctx.JSON(http.StatusOK, dto.PostResponse{
		StatusBool:    true,
		StatusMessage: "Create Post Success..",
		Data:          gin.H{},
	})
}

func (p Post) UpdatePost(ctx *gin.Context) {
	postID, _ := strconv.ParseUint(ctx.Param("postId"), 10, 64)

	var body dto.PostBodyRequest

	if db.Conn.Model(&models.Post{}).Where("id = ?", postID).Find(&body).RowsAffected == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"data": nil, "Message": "Post Not Found .."})
		return
	}

	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	form, _ := ctx.MultipartForm()
	files := form.File["files"]

	if !p.UpdatePostByID(uint(postID), body, files, ctx) {
		ctx.JSON(http.StatusBadRequest, gin.H{"data": nil, "Message": "Update Post Error.."})
		return
	}

	ctx.JSON(http.StatusOK, dto.PostResponse{
		StatusBool:    true,
		StatusMessage: "Update Post Success..",
		Data:          gin.H{},
	})
}

func (p Post) DeletePost(ctx *gin.Context) {
	postID, _ := strconv.ParseUint(ctx.Param("postId"), 10, 64)

	if !p.DeletePostByID(uint(postID)) {
		ctx.JSON(http.StatusOK, gin.H{"data": nil, "Message": "Post Delete Error: "})
		return
	}

	ctx.JSON(http.StatusOK, dto.PostResponse{
		StatusBool:    true,
		StatusMessage: "Post Has Deleted ..",
		Data:          gin.H{"postID": postID},
	})
}
