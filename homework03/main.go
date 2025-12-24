package main

import (
	"encoding/json"
	"fmt"
	homework03 "homework03/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := createDBConnect()
	if err != nil {
		fmt.Println("createDBConnect fail!")
		fmt.Println(err)
	} else {
		fmt.Println("createDBConnect successfully!")
		// 建表
		modelDefinition(db)
		// 新增用户
		addUserData(db)
		// 新增文章
		addPostData(db)
		// 新增评论
		addCommentData(db)
		// 删除评论
		delCommentData(db)
		// 关联查询
		correlatedQuery(db, "zhangsan")

	}

}

// 创建db连接
func createDBConnect() (db *gorm.DB, err error) {

	dsn := "root:xsp888@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Asia%2FShanghai"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})

}

// 模型定义
func modelDefinition(db *gorm.DB) {

	db.AutoMigrate(&homework03.User{}, &homework03.Post{}, &homework03.Comment{})

}

// 新增数据User
func addUserData(db *gorm.DB) {

	users := []*homework03.User{
		{Username: "zhangsan", Email: "111@tt.com", Age: 18},
		{Username: "lisi", Email: "222@tt.com", Age: 19},
		{Username: "wangwu", Email: "333@tt.com", Age: 20},
		{Username: "zhaoliu", Email: "444@tt.com", Age: 21},
		{Username: "tianqi", Email: "555@tt.com", Age: 22},
	}
	db.Create(&users)

}

// 新增数据Comment
func addCommentData(db *gorm.DB) {

	comments := []*homework03.Comment{
		{Content: "good", UserID: 4, PostID: 1},
		{Content: "ok", UserID: 4, PostID: 2},
		{Content: "cha", UserID: 4, PostID: 3},
		{Content: "hao", UserID: 5, PostID: 1},
		{Content: "hao", UserID: 1, PostID: 2},
	}
	db.Create(&comments)

}

// 新增数据Post
func addPostData(db *gorm.DB) {

	posts := []*homework03.Post{
		{Title: "t1", Content: "t1", UserID: 1},
		{Title: "t2", Content: "t2", UserID: 1},
		{Title: "t3", Content: "t3", UserID: 2},
		{Title: "t4", Content: "t4", UserID: 2},
		{Title: "t5", Content: "t5", UserID: 2},
		{Title: "t6", Content: "t6", UserID: 3},
		{Title: "t7", Content: "t7", UserID: 3},
		{Title: "t8", Content: "t8", UserID: 3},
		{Title: "t9", Content: "t9", UserID: 3},
		{Title: "t10", Content: "t10", UserID: 3},
	}
	db.Create(&posts)

}

// 关联查询
func correlatedQuery(db *gorm.DB, name string) {
	// 编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息
	var posts []homework03.Post
	db.Debug().Table("post").
		Select("post.id ID,post.title Title,post.content Content,post.user_id UserID,post.created_at CreatedAt,post.updated_at UpdatedAt").
		Joins("left join user on post.user_id = user.id").
		Where("user.username = ?", name).
		Find(&posts)
	data, _ := json.Marshal(posts)
	fmt.Printf("%s用户对应的文章:%s\n", name, string(data))

	var comments []homework03.Comment
	db.Debug().Table("comment").
		Select("comment.id ID,comment.content Content ,comment.user_id UserID,comment.post_id PostID, comment.created_at CreatedAt,comment.updated_at UpdatedAt").
		Joins("left join user on comment.user_id = user.id").
		Where("user.username = ?", name).
		Find(&comments)
	commentsData, _ := json.Marshal(comments)
	fmt.Printf("%s用户对应的评论:%s", name, string(commentsData))
	//编写Go代码，使用Gorm查询评论数量最多的文章信息
	var postIDs []uint64
	db.Debug().Model(&homework03.Comment{}).
		Select("post_id").
		Group("post_id").
		Having("COUNT(*) = (?)",
			db.Raw("SELECT MAX(cnt) FROM (SELECT COUNT(*) as cnt FROM comment WHERE deleted_at is null GROUP BY post_id) as t"),
		).
		Pluck("post_id", &postIDs)

	fmt.Printf("评论数最大的文章id:%v\n", postIDs)
	var postss []homework03.Post
	db.Debug().Model(&homework03.Post{}).Where("id in ?", postIDs).Find(&postss)
	postssData, _ := json.Marshal(postss)
	fmt.Printf("评论数最多的文章：%s\n", string(postssData))

}

// 删除评论
func delCommentData(db *gorm.DB) {

	var comments []homework03.Comment
	err := db.Where("id in ?", []int{1, 2, 3}).Find(&comments).Error

	if err != nil {
		return
	}
	for _, comment := range comments {
		db.Delete(&comment)
	}

}
