package dto

import "github.com/gin-gonic/gin"

type UserRegisterBodyRequest struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
	Email    string `form:"email" binding:"required"`
}

type UserLoginBodyRequest struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type UserRegisterSuccessResponse struct {
	StatusBool    bool   `json:"StatusBool"`
	StatusMessage string `json:"StatusMessage"`
	Data          gin.H  `json:"Data"`
}

type UserLoginSuccessResponse struct {
	StatusBool    bool   `json:"StatusBool"`
	StatusMessage string `json:"StatusMessage"`
	Data          gin.H  `json:"Data"`
}

type UserFetchAllData struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
