package api

import (
	"Project/dao"
	"Project/model"
	"Project/respond"
	"Project/service"
	"github.com/gin-gonic/gin"
)

// PublishProduction 发表作品
func PublishProduction(c *gin.Context) {
	u := service.GetUsername(c)
	var com model.Comment
	com.Author = u
	com.Writer = ""
	respond.PublishProductionTrue(c)
	dao.BringCommentDate(com)
}
