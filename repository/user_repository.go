package repository

import (
	"github.com/Pugpaprika21/go-gin/db"
	"github.com/Pugpaprika21/go-gin/dto"
	"github.com/Pugpaprika21/go-gin/models"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	CreateUserRegiter(body dto.UserRegisterBodyRequest) bool
	FindUserAll() []dto.UserFetchAllData
	AuthenticateUserJWT(username, password string) (*models.User, error)
}

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		DB: db.Conn,
	}
}

func (u *UserRepository) AuthenticateUserJWT(username, password string) (*models.User, error) {
	var user models.User
	err := u.DB.Where("username = ? AND password = ?", username, password).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepository) CreateUserRegiter(body dto.UserRegisterBodyRequest) bool {
	var user models.User
	if u.DB.Where("username = ?", body.Username).First(&user).RowsAffected > 0 {
		return false
	}
	user = models.User{
		Username: body.Username,
		Password: body.Password,
		Email:    body.Email,
	}
	return u.DB.Create(&user).Error == nil
}

func (u *UserRepository) FindUserAll() []dto.UserFetchAllData {
	var userModel models.User
	var userFetch []dto.UserFetchAllData

	u.DB.Model(&userModel).Find(&userFetch)
	return userFetch
}
