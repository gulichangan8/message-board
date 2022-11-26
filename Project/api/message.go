package api

import (
	"Project/dao"
	"Project/model"
	"Project/respond"
	"Project/service"
	"Project/tool"
	"github.com/gin-gonic/gin"
)

// PublishMessage 发表留言
func PublishMessage(c *gin.Context) {
	u, w, com := service.GetMessage(c)
	Com := tool.CreateMess(w, u, com, "")
	var C model.Use
	C.Mess = Com
	dao.BringDate("message", C)
	respond.PublishMessageTrue(c)
}

// ChangeMessage 修改留言
func ChangeMessage(c *gin.Context) {
	u, m := service.GetNewMessage(c)
	respond.ChangeMessageTrue(c)
	dao.ChangeMessageDate(u, m)
}
