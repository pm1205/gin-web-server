package controller

import (
	"gin-web-server/model"
	"gin-web-server/service"
	"gin-web-server/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Register 用户注册
func Register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:    http.StatusBadRequest,
			Message: "请求参数错误: " + err.Error(),
		})
		return
	}

	if err := service.RegisterUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:    0,
		Message: "注册成功",
		Data: gin.H{
			"userId":   user.ID,
			"username": user.Username,
		},
	})
}

// Login 用户登录
func Login(c *gin.Context) {
	var input model.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{
			Code:    http.StatusBadRequest,
			Message: "请求参数错误: " + err.Error(),
		})
		return
	}

	user, err := service.AuthenticateUser(input.Username, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, model.Response{
			Code:    http.StatusUnauthorized,
			Message: err.Error(),
		})
		return
	}

	token, err := utils.GenerateToken(user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			Code:    http.StatusInternalServerError,
			Message: "生成令牌失败",
		})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Code:    0,
		Message: "登录成功",
		Data: gin.H{
			"token": token,
			"userInfo": gin.H{
				"id":        user.ID,
				"username":  user.Username,
				"avatar":    user.Avatar,
				"createdAt": user.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
				"updatedAt": user.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
			},
		},
	})
}
