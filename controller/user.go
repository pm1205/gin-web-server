package controller

import (
	"net/http"
	"server/model"
	"server/service"
	"server/utils"

	"github.com/gin-gonic/gin"
)

// 注册接口
func Register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code: http.StatusBadRequest,
			Msg:  "Invalid input data",
		})
		return
	}

	if err := service.RegisterUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Code: http.StatusInternalServerError,
			Msg:  "Registration failed: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code: http.StatusOK,
		Msg:  "Registration successful",
	})
}

// 登录接口
func Login(c *gin.Context) {
	var input model.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := service.AuthenticateUser(input.Username, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token generation failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// 获取当前用户信息（需要 JWT 鉴权）
func Profile(c *gin.Context) {
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "This is the protected profile route.",
		"user":    username,
	})
}
