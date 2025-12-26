package main

import (
	"homework04/config"
	controllers "homework04/controller"
	"homework04/middleware"
	"homework04/models"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	// db连接
	db, err := createDBConnect()
	if err != nil {
		log.Printf("createDBConnect fail %s", err)
		// fmt.Println("createDBConnect fail!")
		// fmt.Println(err)
		return
	}
	log.Printf("createDBConnect successfully!")
	// 建表
	modelDefinition(db)

	// 创建Gin实例
	r := gin.Default()

	// 创建控制器实例
	authController := &controllers.AuthController{}
	postController := &controllers.PostController{}
	commentController := &controllers.CommentController{}
	r.GET("/hello", authController.Hello)

	// 公开路由
	public := r.Group("/api/auth")
	{
		public.POST("/register", authController.Register)
		public.POST("/login", authController.Login)

	}

	// 受保护路由（需要JWT认证）
	protected := r.Group("/api")
	protected.Use(middleware.JWTAuth())
	{
		protected.POST("/post/create", postController.Create)
		protected.POST("/post/update", postController.Update)
		protected.POST("/post/delete", postController.Delete)
		protected.POST("/post/getByID", postController.GetByID)
		protected.POST("/post/getAllPost", postController.GetAllPost)
		protected.POST("/comment/create", commentController.Create)
		protected.POST("/comment/getCommentByPostID", commentController.GetCommentByPostID)

	}

	// 启动服务器
	port := ":8080"
	log.Printf("Server starting on port %s", port)
	log.Printf("JWT Secret: %s", config.AppConfig.JWTSecret)
	log.Printf("JWT Expire: %d hours", config.AppConfig.JWTExpire)

	if err := r.Run(port); err != nil {
		log.Fatal("Failed to start server:", err)
	}

}

// 创建db连接
func createDBConnect() (db *gorm.DB, err error) {
	config.InitConfig()
	return config.ConnectDB()

}

// 模型定义
func modelDefinition(db *gorm.DB) {

	db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})

}
