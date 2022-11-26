package api

import "github.com/gin-gonic/gin"

// InterUser 用户接口
func InterUser() {
	r := gin.Default()
	r.Group("/user", func(c *gin.Context) {
		r.POST("/register", Register)
		r.POST("/login", Login)
		r.POST("/question", Question)
		r.POST("/answer_question", AnswerQuestion)
		r.PUT("/change_password", ChangePassword)
	})
	err := r.Run()
	if err != nil {
		return
	}
}

// InterMessage 留言接口
func InterMessage() {
	r := gin.Default()
	r.Group("message", func(c *gin.Context) {
		r.POST("/publish_comment", PublishMessage)
		r.PUT("/change_message", ChangeMessage)
	})
	err := r.Run()
	if err != nil {
		return
	}
}
