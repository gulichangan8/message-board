package api

import (
	"github.com/gin-gonic/gin"
)

func InitEngine() {
	// 用户接口
	r := gin.Default()
	user := r.Group("/user")
	{
		user.POST("/register", Register)
		user.POST("/login", Login)
		user.POST("/question", Question)
		user.POST("/answer_question", AnswerQuestion)
		user.PUT("/change_password", ChangePassword)
		user.POST("/personal_message", PersonalMessage)
		user.PUT("/change_personal_message", ChangePersonalMessage)
		user.DELETE("logout", DeleteLogin)
	}

	// 留言接口
	message := r.Group("/message")
	{
		message.POST("/publish", PublishMessage)
		message.PUT("/change", ChangeMessage)
		message.DELETE("/delete", DeleteMessage)
		message.POST("/respond", RespondMessage)
		message.GET("/read", ReadMessage)
	}

	// 发布作品接口
	r.POST("/publish_production", PublishProduction)

	// 评论作品接口
	comment := r.Group("/comment")
	{
		//以下是原表格（comment）评论（懒得删了,函数一层一层的太多啦）
		comment.POST("/publish", PublishComment)
		comment.PUT("/change", ChangeComment)
		comment.DELETE("/delete", DeleteComment)
		comment.POST("noname_comment", NoNameComment)
		//这两个不受comment影响
		comment.PUT("/good", ChangeGood)
		comment.GET("good_member", GetGoodMember)
		//以下是comments嵌套评论
		//comment.POST("/publish", PublishComment)
		//comment.GET("comments", GetAuthorComment)
		//comment.POST("publish_comments", PublishComments)
		//comment.GET("my_comments", GetMyComment)
		//comment.PUT("/change", ChangeComments)
		//comment.DELETE("/delete", DeleteComments)
	}
	err := r.Run()
	if err != nil {
		return
	}
}
