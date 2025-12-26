package controller

import (
	"homework04/models"
	services "homework04/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct{}

func (ac *AuthController) Register(c *gin.Context) {
	var registerReq models.User

	// 绑定JSON数据
	if err := c.ShouldBindJSON(&registerReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "请求参数错误",
			"details": err.Error(),
		})
		return
	}

	// 调用注册服务
	user, err := services.Register(registerReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "注册成功",
		"user":    user,
	})
}

func (ac *AuthController) Hello(c *gin.Context) {

	c.JSON(http.StatusCreated, gin.H{
		"message": "hello",
	})
}

func (ac *AuthController) Login(c *gin.Context) {
	var loginReq models.User

	// 绑定JSON数据
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "请求参数错误",
			"details": err.Error(),
		})
		return
	}

	// 调用登录服务
	token, user, err := services.Login(loginReq)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "登录成功",
		"token":      token,
		"user":       user,
		"expires_in": 24 * 60 * 60, // token过期时间（秒）
	})
}
