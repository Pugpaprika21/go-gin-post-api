package controller

import (
	"net/http"
	"os"
	"time"

	"github.com/Pugpaprika21/go-gin/dto"
	"github.com/Pugpaprika21/go-gin/models"
	"github.com/Pugpaprika21/go-gin/repository"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	repository.UserRepositoryInterface
}

var secretKey = []byte(os.Getenv("APP_TOKEN"))

func (u User) GetUserAll(ctx *gin.Context) {
	users := u.FindUserAll()
	if len(users) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"Data": nil, "Message": "Login error.."})
		return
	}
	ctx.JSON(http.StatusOK, dto.UserLoginSuccessResponse{
		StatusBool:    true,
		StatusMessage: "Login Success..",
		Data:          gin.H{"users": users},
	})
}

func (u User) Register(ctx *gin.Context) {
	var body dto.UserRegisterBodyRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	success := u.CreateUserRegiter(body)
	if !success {
		ctx.JSON(http.StatusBadRequest, gin.H{"Data": nil, "Message": "Register error.."})
		return
	}
	ctx.JSON(http.StatusOK, dto.UserRegisterSuccessResponse{
		StatusBool:    true,
		StatusMessage: "Register Success..",
		Data:          gin.H{},
	})
}

func (u User) Login(ctx *gin.Context) {
	var body dto.UserLoginBodyRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := u.AuthenticateUserJWT(body.Username, body.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Data": nil, "Message": "Login error.."})
		return
	}

	token, _ := u.generateJWTToken(user)

	ctx.JSON(http.StatusOK, dto.UserLoginSuccessResponse{
		StatusBool:    true,
		StatusMessage: "Login Success..",
		Data: gin.H{
			"Token": token,
		},
	})
}

func (u User) generateJWTToken(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"data": gin.H{
			"user": user,
		},
		"exp": time.Now().Add(time.Minute).Unix(),
	})

	tokenJWT, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenJWT, nil
}
