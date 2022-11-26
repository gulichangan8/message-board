package api

import (
	"Project/dao"
	"Project/model"
	"Project/respond"
	"Project/service"
	"github.com/gin-gonic/gin"
)

func PublishProduction(c *gin.Context) {
	u := service.GetUsername(c)
	var use model.Use
	use.Comment.Author = u
	use.Comment.WriterAndMessage = ""
	respond.PublishProductionTrue(c)
	dao.BringDate("comment", use)
}
