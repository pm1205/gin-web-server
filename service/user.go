package service

import (
	"errors"

	"gin-web-server/config"
	"gin-web-server/model"

	"golang.org/x/crypto/bcrypt"
)

// RegisterUser 处理用户注册逻辑
func RegisterUser(user *model.User) error {
	// 1. 检查用户名是否已存在
	var existing model.User
	if err := config.DB.Where("username = ?", user.Username).First(&existing).Error; err == nil {
		return errors.New("用户名已存在")
	}

	// 2. 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("密码加密失败")
	}
	user.Password = string(hashedPassword)

	// 3. 设置默认头像
	if user.Avatar == "" {
		user.Avatar = "https://api.dicebear.com/7.x/avataaars/svg?seed=" + user.Username
	}

	// 4. 插入数据库
	if err := config.DB.Create(user).Error; err != nil {
		return errors.New("创建用户失败: " + err.Error())
	}

	return nil
}

// AuthenticateUser 处理用户登录验证
func AuthenticateUser(username, password string) (*model.User, error) {
	var user model.User
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, errors.New("用户不存在")
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("密码错误")
	}

	// 检查用户状态
	if user.Status != 1 {
		return nil, errors.New("账号已被禁用")
	}

	return &user, nil
}

// GetUserByID 根据ID获取用户信息
func GetUserByID(id uint) (*model.User, error) {
	var user model.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return nil, errors.New("用户不存在")
	}
	return &user, nil
}

// GetUserByUsername 根据用户名获取用户信息
func GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, errors.New("用户不存在")
	}
	return &user, nil
}
