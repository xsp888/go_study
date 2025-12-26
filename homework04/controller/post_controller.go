package controller

import (
	"homework04/models"
	"homework04/request"
	"homework04/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostController struct{}

func (pc *PostController) Create(c *gin.Context) {
	var postReq models.Post
	// 绑定JSON数据
	if err := c.ShouldBindJSON(&postReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "请求参数错误",
			"details": err.Error(),
		})
		return
	}
	userID := c.GetUint("user_id")
	if err := service.NewPostService().Create(postReq, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "程序异常",
			"details": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"error":   "",
			"details": "成功",
		})
	}

}

func (pc *PostController) GetByID(c *gin.Context) {
	var postReq models.Post
	// 绑定JSON数据
	if err := c.ShouldBindJSON(&postReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "请求参数错误",
			"details": err.Error(),
		})
		return
	}
	post, err := service.NewPostService().GetByID(postReq.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "程序异常",
			"details": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, post)
	}

}

func (pc *PostController) GetAllPost(c *gin.Context) {
	var postRequest request.PostRequest
	// 绑定JSON数据
	if err := c.ShouldBindJSON(&postRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "请求参数错误",
			"details": err.Error(),
		})
		return
	}

	userID := c.GetUint("user_id")
	postResp, err := service.NewPostService().GetAllPost(postRequest.Page, postRequest.PageSize, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "程序异常",
			"details": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, postResp)
	}

}

func (pc *PostController) Update(c *gin.Context) {
	var postReq models.Post
	// 绑定JSON数据
	if err := c.ShouldBindJSON(&postReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "请求参数错误",
			"details": err.Error(),
		})
		return
	}
	userID := c.GetUint("user_id")
	if err := service.NewPostService().Update(postReq.ID, postReq, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "程序异常",
			"details": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"error":   "",
			"details": "成功",
		})
	}

}

func (pc *PostController) Delete(c *gin.Context) {
	var postReq models.Post
	// 绑定JSON数据
	if err := c.ShouldBindJSON(&postReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "请求参数错误",
			"details": err.Error(),
		})
		return
	}
	userID := c.GetUint("user_id")
	if err := service.NewPostService().Delete(postReq.ID, userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "程序异常",
			"details": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"error":   "",
			"details": "成功",
		})
	}

}
