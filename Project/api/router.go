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
}

// InterComment 评论作品接口
func InterComment() {
	r := gin.Default()
	r.Group("/comment", func(c *gin.Context) {
		r.POST("/publish", PublishComment)
		r.PUT("/change", ChangeComment)
		r.DELETE("/delete", DeleteComment)
	})
	err := r.Run()
	if err != nil {
		return
	}
}
