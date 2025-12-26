package service

import (
	"errors"
	"homework04/config"
	"homework04/models"
	"homework04/utils"

	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Register(userReq models.User) (*models.User, error) {
	// 检查用户名是否已存在
	var existingUser models.User
	result := config.DB.Where("username = ?", userReq.Username).First(&existingUser)
	if result.RowsAffected > 0 {
		return nil, errors.New("用户名已存在")
	}

	// 检查邮箱是否已存在
	result = config.DB.Where("email = ?", userReq.Email).First(&existingUser)
	if result.RowsAffected > 0 {
		return nil, errors.New("邮箱已存在")
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// 创建用户
	user := models.User{
		Username: userReq.Username,
		Password: string(hashedPassword),
		Email:    userReq.Email,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	// 返回用户信息（不包含密码）
	userResponse := &models.User{
		Model: gorm.Model{
			ID: user.ID,
		},
		Username: user.Username,
		Email:    user.Email,
	}

	return userResponse, nil
}

func Login(loginReq models.User) (string, *models.User, error) {
	var user models.User

	// 查找用户
	result := config.DB.Where("username = ?", loginReq.Username).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return "", nil, errors.New("用户不存在")
	}

	// 验证密码
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password))
	if err != nil {
		return "", nil, errors.New("密码错误")
	}

	// 生成JWT token
	token, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		return "", nil, err
	}
	log.Printf("Username is %s, token is %s", user.Username, token)
	// 返回用户信息
	userResponse := &models.User{
		Model: gorm.Model{
			ID: user.ID,
		},

		Username: user.Username,
		Email:    user.Email,
	}

	return token, userResponse, nil
}
