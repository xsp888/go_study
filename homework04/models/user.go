package models

import (
	"gorm.io/gorm"
)

// 用户
type User struct {
	gorm.Model
	Username string `gorm:"size:50;index;not null;comment:用户名称" json:"username"`
	Password string `gorm:"size:500;not null;comment:用户密码" json:"password"`
	Email    string `gorm:"size:100;index;not null;comment:邮箱" json:"email"`
	Age      uint8  `gorm:"index;not null;comment:年龄" json:"age"`
	PostNum  uint64 `gorm:"comment:文章数量" json:"post_num"`

	// 一对多关系：一个用户有多篇文章
	Posts []Post `gorm:"foreignKey:UserID;" json:"posts"`
}

func (User) TableName() string {
	return "user"
}

func (User) TableComment() string {
	return "用户信息表"
}
