package models

import "gorm.io/gorm"

// 评论
type Comment struct {
	gorm.Model
	Content string `gorm:"type:text;not null;comment:信息" json:"content"`
	// 外键：评论者
	UserID uint `gorm:"not null;index;comment:用户id" json:"user_id"`
	// 外键：所属文章
	PostID uint `gorm:"not null;index;comment:文章id" json:"post_id"`

	// 属于关系：评论属于用户和文章
	User User `gorm:"foreignKey:UserID" json:"user"`
	Post Post `gorm:"foreignKey:PostID" json:"post"`
}

func (Comment) TableName() string {
	return "comment"
}

func (Comment) TableComment() string {
	return "评论表"
}

// 为 Comment 模型添加一个钩子函数，在评论增加时自动增加文章的评论数量
func (c *Comment) AfterCreate(tx *gorm.DB) error {
	err := tx.Debug().Model(&Post{}).Where("id = ?", c.PostID).Update("comment_num", gorm.Expr("comment_num + ?", 1)).Error

	return err

}

// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，
// 如果评论数量为 0，则更新文章的评论状态为 "无评论"
// 竟然不支持批量删除？？？？？？
func (c *Comment) AfterDelete(tx *gorm.DB) error {

	var commentNum uint64
	tx.Debug().Model(&Post{}).
		Where("id = ?", c.PostID).
		Select("comment_num").
		Pluck("comment_num", &commentNum)
	var err error
	if commentNum-1 == 0 {
		err = tx.Debug().Model(&Post{}).Where("id = ?", c.PostID).Update("comment_status", "无评论").Error
		if err != nil {
			return err
		}

	}
	err = tx.Debug().Model(&Post{}).Where("id = ?", c.PostID).Update("comment_num", gorm.Expr("comment_num - ?", 1)).Error

	return err

}
