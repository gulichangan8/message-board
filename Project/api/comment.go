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
	a, w, com := service.GetUsernameCommentWriteTo(c)
	ok := service.CheckAuthorExist(w)
	if ok {
		C := tool.CreateComment(a, w, com)
		var use model.Use
		use.Comment = C
		dao.BringDate("comment", use)
		respond.PublishCommentTrue(c)
	} else {
		respond.PublishCommentErr(c)
	}
}

// ChangeComment 修改评论
func ChangeComment(c *gin.Context) {
	a, w, com := service.GetUsernameCommentWriteTo(c)
	ok := service.CheckAuthorExist(a) && service.CheckWriteExist(w)
	if ok {
		dao.ChangeCommentDate(com, a)
		respond.ChangeCommentTrue(c)
	} else {
		respond.ChangeCommentErr(c)
	}
}

// DeleteComment 删除评论
func DeleteComment(c *gin.Context) {
	w := service.GetUsername(c)
	ok := service.CheckAuthorExist(w)
	if ok {
		dao.ChangeCommentDate("", w)
		respond.ChangeCommentTrue(c)
	} else {
		respond.ChangeCommentErr(c)
	}
}
