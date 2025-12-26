package models

import (
	"gorm.io/gorm"
)

// 文章
type Post struct {
	gorm.Model
	Title   string `gorm:"size:200;not null;comment:标题" json:"title"`
	Content string `gorm:"type:text;not null;comment:信息" json:"content"`
	// 外键
	UserID        uint   `gorm:"not null;index;comment:用户id" json:"user_id"`
	CommentStatus string `gorm:"size:200;comment:评论状态" json:"comment_status"`
	CommentNum    uint64 `gorm:"comment:评论数量" json:"comment_num"`
	// 属于关系：文章属于用户
	User User `gorm:"foreignKey:UserID" json:"user"`

	// 一对多关系：一篇文章有多个评论
	Comments []Comment `gorm:"foreignKey:PostID;" json:"comments"`
}

func (Post) TableName() string {
	return "post"
}

func (Post) TableComment() string {
	return "文章表"
}

// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段
func (p *Post) AfterCreate(tx *gorm.DB) error {
	err := tx.Debug().Model(&User{}).Where("id = ?", p.UserID).Update("post_num", gorm.Expr("post_num + ?", 1)).Error

	return err

}
