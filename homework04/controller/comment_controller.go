package controller

import (
	"homework04/models"
	"homework04/request"
	"homework04/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommentController struct{}

func (pc *CommentController) Create(c *gin.Context) {
	var commentReq models.Comment
	// 绑定JSON数据
	if err := c.ShouldBindJSON(&commentReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "请求参数错误",
			"details": err.Error(),
		})
		return
	}
	userID := c.GetUint("user_id")
	if err := service.NewCommentService().Create(commentReq, userID); err != nil {
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

func (pc *CommentController) GetCommentByPostID(c *gin.Context) {
	var commentRequest request.CommentRequest
	// 绑定JSON数据
	if err := c.ShouldBindJSON(&commentRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "请求参数错误",
			"details": err.Error(),
		})
		return
	}

	postResp, err := service.NewCommentService().GetCommentByPostID(commentRequest.Page, commentRequest.PageSize, commentRequest.PostID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "程序异常",
			"details": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, postResp)
	}

}
