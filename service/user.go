package service

import (
	"errors"

	"server/config"
	"server/model"

	"golang.org/x/crypto/bcrypt"
)

// RegisterUser 处理用户注册逻辑
func RegisterUser(user *model.User) error {
	// 1. 检查用户名是否已存在
	var existing model.User
	if err := config.DB.Where("username = ?", user.Username).First(&existing).Error; err == nil {
		return errors.New("username already exists")
	}

	// 2. 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password")
	}
	user.Password = string(hashedPassword)

	// 3. 插入数据库
	if err := config.DB.Create(user).Error; err != nil {
		return errors.New("failed to create user: " + err.Error())
	}

	return nil
}

func AuthenticateUser(username, password string) (*model.User, error) {
	var user model.User
	if err := config.DB.Where("username = ? AND password = ?", username, password).First(&user).Error; err != nil {
		return nil, errors.New("invalid credentials")
	}
	return &user, nil
}
