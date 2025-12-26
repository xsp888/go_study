package service

import (
	"errors"
	"fmt"
	"homework04/config"
	"homework04/models"
	"homework04/response"
	"time"

	"gorm.io/gorm"
)

type PostService struct{}

func NewPostService() *PostService {
	return &PostService{}
}

// 创建文章
func (s *PostService) Create(req models.Post, userID uint) error {

	var user models.User
	err := config.DB.First(&user, userID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("%d 作者未认证", userID)
		}
		return err
	}
	if len(req.Title) == 0 || len(req.Content) == 0 {
		return fmt.Errorf("文章的标题和内容未传")

	}
	// 创建文章
	post := models.Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  userID,
	}

	if err := config.DB.Create(&post).Error; err != nil {

		return fmt.Errorf("创建文章失败: %w", err)
	}
	return nil

}

// 获取单个文章
func (s *PostService) GetByID(postID uint) (*models.Post, error) {
	var post models.Post

	err := config.DB.First(&post, postID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("文章不存在")
		}
		return nil, err
	}

	return &models.Post{
		Model: gorm.Model{
			ID:        post.ID,
			CreatedAt: post.CreatedAt,
			UpdatedAt: post.UpdatedAt,
		},
		Title:   post.Title,
		Content: post.Content,
		UserID:  post.UserID,
	}, nil
}

// 获取所有文章（分页）
func (s *PostService) GetAllPost(page, pageSize int, userID uint) (*response.PostResponse, error) {
	var posts []models.Post
	var total int64

	// 计算偏移量
	offset := (page - 1) * pageSize
	if offset < 0 {
		offset = 0
	}

	// 获取总数
	config.DB.Model(&models.Post{}).Where("user_id = ?", userID).Count(&total)

	// 获取分页数据
	result := config.DB.Where("user_id = ?", userID).
		Order("created_at DESC").
		Limit(pageSize).
		Offset(offset).
		Find(&posts)

	if result.Error != nil {
		return nil, result.Error
	}

	// 构建响应列表
	return &response.PostResponse{
		Posts:    posts,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}, nil
}

// 更新文章（仅作者可更新）
func (s *PostService) Update(postID uint, req models.Post, userID uint) error {
	// 首先检查文章是否为该作者
	var post models.Post
	result := config.DB.First(&post, postID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("未找到该作者的文章不存在")
		}
		return result.Error
	}
	if post.UserID != userID {
		return errors.New("没有权限更新此文章")
	}
	if len(req.Title) == 0 && len(req.Content) == 0 {
		return errors.New("标题与内容不能同时为空")
	}

	// 更新字段
	updates := make(map[string]interface{})
	if len(req.Title) != 0 {
		updates["title"] = req.Title
	}
	if len(req.Content) != 0 {
		updates["content"] = req.Content
	}
	if len(updates) != 0 {
		updates["updated_at"] = time.Now().Truncate(time.Millisecond)
	}

	// 执行更新

	if err := config.DB.Model(&models.Post{}).Where("id = ?", postID).Updates(updates).Error; err != nil {
		return fmt.Errorf("更新文章失败: %w", err)
	}

	return nil
}

// 删除文章（仅作者可删除）
func (s *PostService) Delete(postID uint, userID uint) error {
	// 首先检查文章是否为该作者
	var post models.Post
	result := config.DB.First(&post, postID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("未找到该作者的文章不存在")
		}
		return result.Error
	}
	if post.UserID != userID {
		return errors.New("没有权限更新此文章")
	}

	// 删除文章
	if err := config.DB.Delete(&post).Error; err != nil {
		return fmt.Errorf("删除文章失败: %w", err)
	}

	return nil
}
