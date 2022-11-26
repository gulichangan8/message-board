package api

import (
	"Project/dao"
	"Project/model"
	"Project/respond"
	"Project/service"
	"Project/tool"
	"github.com/gin-gonic/gin"
)

// PublishComment 发表评论
func PublishComment(c *gin.Context) {
	w, com := service.GetUsernameCommentWriteTo(c)
	ok := service.CheckAuthorExist(w)
	if ok {
		C := tool.CreateComment(w, com)
		var use model.Use
		use.Comment = C
		dao.BringDate("comment", use)
		respond.PublishCommentTrue(c)
	} else {
		respond.PublishCommentErr(c)
	}
}
