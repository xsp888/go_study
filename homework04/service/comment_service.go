package service

import (
	"errors"
	"fmt"
	"homework04/config"
	"homework04/models"
	"homework04/response"

	"gorm.io/gorm"
)

type CommentService struct{}

func NewCommentService() *CommentService {
	return &CommentService{}
}

// 创建评论
func (s *CommentService) Create(req models.Comment, userID uint) error {

	var user models.User
	err := config.DB.First(&user, userID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("%d 作者未认证", userID)
		}
		return err
	}
	if len(req.Content) == 0 {
		return fmt.Errorf("评论内容未传")

	}
	// 创建评论
	comment := models.Comment{
		Content: req.Content,
		UserID:  userID,
		PostID:  req.PostID,
	}

	if err := config.DB.Create(&comment).Error; err != nil {

		return fmt.Errorf("创建评论失败: %w", err)
	}
	return nil

}

// 根据文章ID查询所有评论
func (s *CommentService) GetCommentByPostID(page, pageSize int, postID uint) (*response.CommentResponse, error) {
	var comments []models.Comment
	var total int64

	// 计算偏移量
	offset := (page - 1) * pageSize
	if offset < 0 {
		offset = 0
	}

	// 获取总数
	config.DB.Model(&models.Comment{}).Where("post_id = ?", postID).Count(&total)

	// 获取分页数据
	result := config.DB.Where("post_id = ?", postID).
		Order("created_at DESC").
		Limit(pageSize).
		Offset(offset).
		Find(&comments)

	if result.Error != nil {
		return nil, result.Error
	}

	// 构建响应列表
	return &response.CommentResponse{
		Comments: comments,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}, nil
}
