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
