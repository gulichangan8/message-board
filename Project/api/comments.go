package api

import (
	"Project/dao"
	"Project/respond"
	"Project/service"
	"github.com/gin-gonic/gin"
)

// GetAuthorComment 获得相应作者下的全部评论
func GetAuthorComment(c *gin.Context) {
	bi := service.GetAuthorComment(c)
	respond.GetAuthorCommentTrue(c, bi)
}

// PublishComments 发表评论
func PublishComments(c *gin.Context) {
	parentId, author, writer, comment := service.GetCom(c)
	dao.BringComments(parentId, author, writer, comment)
	respond.PublishCommentTrue(c)
}

// GetMyComment 查看自己发表过的评论
func GetMyComment(c *gin.Context) {
	com := service.GetWriterCom(c)
	respond.GetMyCommentTrue(c, com)
}

// ChangeComments 修改评论
func ChangeComments(c *gin.Context) {
	id, com := service.GetWriterId(c)
	dao.ChangeComments(id, com)
	respond.ChangeCommentTrue(c)
}
