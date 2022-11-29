package api

import (
	"github.com/gin-gonic/gin"
)

// InterUser 用户接口
func InterUser() {
	r := gin.Default()
	r.Group("/user", func(c *gin.Context) {
		r.POST("/register", Register)
		r.POST("/login", Login)
		r.POST("/question", Question)
		r.POST("/answer_question", AnswerQuestion)
		r.PUT("/change_password", ChangePassword)
		r.POST("/personal_message", PersonalMessage)
		r.PUT("/change_personal_message", ChangePersonalMessage)
		r.DELETE("logout", DeleteLogin)
	})
	err := r.Run()
	if err != nil {
		return
	}
}

// InterMessage 留言接口
func InterMessage() {
	r := gin.Default()
	r.Group("/message", func(c *gin.Context) {
		r.POST("/publish", PublishMessage)
		r.PUT("/change", ChangeMessage)
		r.DELETE("/delete", DeleteMessage)
		r.POST("/respond", RespondMessage)
		r.GET("/read", ReadMessage)
	})
	err := r.Run()
	if err != nil {
		return
	}
}

// InterPublishProduction 发布作品接口
func InterPublishProduction() {
	r := gin.Default()
	r.POST("/publish_production", PublishProduction)
	err := r.Run()
	if err != nil {
		return
	}
}

// InterComment 评论作品接口
func InterComment() {
	r := gin.Default()
	r.Group("/comment", func(c *gin.Context) {
		//以下是原表格（comment）评论（懒得删了,函数一层一层的太多啦）
		r.POST("/publish", PublishComment)
		r.PUT("/change", ChangeComment)
		r.DELETE("/delete", DeleteComment)
		r.POST("noname_comment", NoNameComment)
		//这两个不受comment影响
		r.PUT("/good", ChangeGood)
		r.GET("good_member", GetGoodMember)
		//以下是comments嵌套评论
		r.POST("/publish", PublishComment)
		r.GET("comments", GetAuthorComment)
		r.POST("publish_comments", PublishComments)
		r.GET("my_comments", GetMyComment)
		r.PUT("/change", ChangeComments)
	})
	err := r.Run()
	if err != nil {
		return
	}
}
