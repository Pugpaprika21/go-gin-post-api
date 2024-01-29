package controller

import (
	"net/http"
	"strconv"

	"github.com/Pugpaprika21/go-gin/dto"
	"github.com/Pugpaprika21/go-gin/repository"
	"github.com/gin-gonic/gin"
)

type Comment struct {
	repository.CommentRepositoryInterface
}

func (c Comment) CreateComment(ctx *gin.Context) {
	var body dto.CommentBodyRequest
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !c.CreateUserComment(body) {
		ctx.JSON(http.StatusBadRequest, gin.H{"Data": nil, "Message": "Comment Not Found Or Comment Create Error.."})
		return
	}

	ctx.JSON(http.StatusOK, dto.CommentResponse{
		StatusBool:    true,
		StatusMessage: "Comment Create Success ..",
		Data:          gin.H{"Comment": body},
	})
}

func (c Comment) GetCommentAll(ctx *gin.Context) {
	comments := c.FindAllComment()
	if len(comments) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"Data": nil, "Message": "Comment Is Empty .."})
		return
	}

	ctx.JSON(http.StatusOK, dto.CommentResponse{
		StatusBool:    true,
		StatusMessage: "Comment ..",
		Data:          gin.H{"comments": comments},
	})
}

func (c Comment) GetComment(ctx *gin.Context) {
	actionParam := ctx.Param("action")
	commentID, _ := strconv.ParseUint(ctx.Param("commentId"), 10, 64)
	userID, _ := strconv.ParseUint(ctx.Param("userId"), 10, 64)

	if actionParam == "_comment" {
		data := c.FindCommentByID(uint(commentID))
		ctx.JSON(http.StatusOK, dto.CommentResponse{
			StatusBool:    true,
			StatusMessage: "Comment ..",
			Data:          gin.H{"comments": data},
		})
	} else if actionParam == "_user" {
		data := c.FindCommentByUserID(uint(userID))
		ctx.JSON(http.StatusOK, dto.CommentResponse{
			StatusBool:    true,
			StatusMessage: "Comment By User ..",
			Data:          gin.H{"comments": data},
		})
	}
}

func (c Comment) UpdateComment(ctx *gin.Context) {
	commentID, _ := strconv.ParseUint(ctx.Param("commentId"), 10, 64)
	var body dto.CommentBodyRequest

	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !c.UpdateCommentByID(uint(commentID), body) {
		ctx.JSON(http.StatusBadRequest, gin.H{"Data": nil, "Message": "Update Comment Error .."})
		return
	}

	ctx.JSON(http.StatusOK, dto.CommentResponse{
		StatusBool:    true,
		StatusMessage: "Comment Update ..",
		Data:          gin.H{},
	})
}

func (c Comment) DeleteComment(ctx *gin.Context) {
	commentID, _ := strconv.ParseUint(ctx.Param("commentId"), 10, 64)

	if !c.DeleteCommentByID(uint(commentID)) {
		ctx.JSON(http.StatusBadRequest, gin.H{"Data": nil, "Message": "Delete Comment Error .."})
		return
	}
	ctx.JSON(http.StatusOK, dto.CommentResponse{
		StatusBool:    true,
		StatusMessage: "Comment Delete ..",
		Data:          gin.H{},
	})
}
